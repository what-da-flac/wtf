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
	queueName := string(golang.MediaProcess)
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
		srcPathName := config.Paths.Resolve(msg.SrcPathName)
		if srcPathName == "" {
			err := fmt.Errorf("invalid path name: %s", msg.SrcPathName)
			logger.Error(err)
			return true, err
		}
		src := filepath.Join(srcPathName, msg.Filename)
		dstPathName := config.Paths.Resolve(msg.DstPathName)
		if dstPathName == "" {
			err := fmt.Errorf("invalid path name: %s", msg.DstPathName)
			logger.Error(err)
			return true, err
		}
		dst := filepath.Join(dstPathName, msg.Filename) + ".m4a"
		// on any case, original file will be deleted
		defer func() {
			_ = os.Remove(src)
		}()
		audio, err := ExtractAudio(src)
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
		bitRate = CalculateNumber(bitRate, minBitRate)
		// convert file to desired format/resolution
		if err = commands.CmdFFMpegAudio(src, dst, bitRate); err != nil {
			logger.Error(err)
			return true, err
		}
		// TODO: write final audio file to db
		logger.Infof("ready: %s source bit rate:%d destination bit rate: %d content type: %s",
			msg.OriginalFilename, audio.BitRate, bitRate, msg.DestinationContentType)
		return true, nil
	}
}

func HasAudioEnoughQuality(bitRate, minBitrate int) bool {
	return bitRate >= minBitrate
}

func CalculateNumber(current, dst int) int {
	if current < dst {
		return current
	}
	return dst
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
