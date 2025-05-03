package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/what-da-flac/wtf/go-common/brokers"
	"github.com/what-da-flac/wtf/go-common/commands"
	"github.com/what-da-flac/wtf/go-common/env"
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
	return serve(logger)
}

func serve(zl *zap.Logger) error {
	logger := zl.Sugar()
	ctx := context.Background()
	config := env.New()
	// TODO: set redis connection from environment variables
	client := brokers.NewClient()
	queueName := string(golang.QueueNameMediainfo)
	subscriber, err := brokers.NewSubscriber[golang.MediaInfoInput](client, queueName, "media-info")
	if err != nil {
		return err
	}
	processMessageFn := func(msg golang.MediaInfoInput) (ack bool, err error) {
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
		if !HasAudioEnoughQuality(*audio, msg.MinBitrate) {
			err := fmt.Errorf("audio bitdepth: %d is less than minimum: %d", audio.SamplingRate, msg.MinBitrate)
			logger.Error(err)
			return true, err
		}
		// TODO: write final audio file to db
		logger.Infof("audio ready to be processed: %d", audio.SamplingRate)
		return true, nil
	}
	errFn := func(err error) {
		log.Fatal(err)
	}
	zl.Sugar().Infoln("starting subscriber:", queueName)
	subscriber.Listen(ctx, processMessageFn, errFn)
	return nil
}

func HasAudioEnoughQuality(audio golang.Audio, minBitrate int) bool {
	return audio.SamplingRate >= minBitrate
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
