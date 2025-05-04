package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/what-da-flac/wtf/go-common/brokers"
	"github.com/what-da-flac/wtf/go-common/commands"
	"github.com/what-da-flac/wtf/go-common/env"
	"github.com/what-da-flac/wtf/go-common/ifaces"
	"github.com/what-da-flac/wtf/openapi/domains"
	"github.com/what-da-flac/wtf/openapi/gen/golang"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	logger, err := zap.NewProductionConfig().Build()
	if err != nil {
		return err
	}
	config := env.New()
	return listen(config, logger.Sugar())
}

func listen(config *env.Config, logger ifaces.Logger) error {
	ctx := context.Background()
	// TODO: set redis connection from environment variables
	client := brokers.NewClient()
	queueName := string(golang.QueueNameMediaProcess)
	subscriber, err := brokers.NewSubscriber[golang.MediaInfoInput](client, queueName, queueName)
	if err != nil {
		return err
	}
	errFn := func(err error) {
		logger.Error(err)
	}
	logger.Info("starting subscriber:", queueName)
	subscriber.Listen(ctx, processMessageFn(config, logger), errFn)
	return nil
}

func processMessageFn(config *env.Config, logger ifaces.Logger) func(msg golang.MediaInfoInput) (ack bool, err error) {
	return func(msg golang.MediaInfoInput) (ack bool, err error) {
		pathName := config.Paths.Resolve(msg.PathName)
		if pathName == "" {
			err := fmt.Errorf("invalid path name: %s", msg.PathName)
			logger.Error(err)
			return true, err
		}
		filename := filepath.Join(pathName, msg.Filename)
		// on any case, original file will be deleted
		defer func() {
			_ = os.Remove(filename)
		}()
		audio, err := ExtractAudio(filename)
		if err != nil {
			logger.Error(err)
			return true, err
		}
		bitRate := audio.BitRate
		minBitRate := msg.MinBitrate
		if !HasAudioEnoughQuality(bitRate, minBitRate) {
			err := fmt.Errorf("audio bitdepth: %d is less than minimum: %d", bitRate, minBitRate)
			logger.Error(err)
			return true, err
		}
		// determine bitrate and convert audio file to m4a
		bitRate = CalculateBitrate(bitRate, minBitRate)
		// TODO: write final audio file to db
		logger.Infof("ready: %s source bit rate:%d destination bit rate: %d content type: %s",
			msg.OriginalFilename, audio.BitRate, bitRate, msg.DestinationContentType)
		return true, nil
	}
}

func HasAudioEnoughQuality(bitRate, minBitrate int) bool {
	return bitRate >= minBitrate
}

func CalculateBitrate(bitRate, dstBitRate int) int {
	if bitRate < dstBitRate {
		return bitRate
	}
	return dstBitRate
}

func ExtractAudio(filename string) (*golang.Audio, error) {
	// read mediainfo
	data, err := commands.CmdMediaInfo(filename)
	if err != nil {
		return nil, err
	}
	info, err := domains.NewMediaInfo(data)
	if err != nil {
		return nil, err
	}
	audio := domains.NewAudio(info)
	return &audio, nil
}
