package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	_ "github.com/lib/pq"
	"github.com/what-da-flac/wtf/go-common/brokers"
	"github.com/what-da-flac/wtf/go-common/commands"
	"github.com/what-da-flac/wtf/go-common/env"
	"github.com/what-da-flac/wtf/go-common/identifiers"
	"github.com/what-da-flac/wtf/go-common/ifaces"
	"github.com/what-da-flac/wtf/go-common/pgpq"
	"github.com/what-da-flac/wtf/go-common/repositories"
	"github.com/what-da-flac/wtf/openapi/domains"
	"github.com/what-da-flac/wtf/openapi/gen/golang"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

type Container struct {
	config     *env.Config
	identifier ifaces.Identifier
	logger     ifaces.Logger
	repository ifaces.Repository
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	config := env.New()
	logger, err := zap.NewProductionConfig().Build()
	if err != nil {
		return err
	}
	db, err := pgpq.New(config.DB.URL)
	if err != nil {
		return err
	}
	repository, err := repositories.NewPgRepo(db, config.DB.URL, false)
	if err != nil {
		return err
	}
	container := &Container{
		config:     config,
		identifier: identifiers.NewIdentifier(),
		repository: repository,
		logger:     logger.Sugar(),
	}
	return listen(container)
}

func listen(container *Container) error {
	logger := container.logger
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
	subscriber.Listen(ctx, processMessageFn(container), errFn)
	return nil
}

func processMessageFn(container *Container) func(msg golang.MediaInfoInput) (ack bool, err error) {
	config := container.config
	logger := container.logger
	return func(msg golang.MediaInfoInput) (ack bool, err error) {
		srcPathName := config.Paths.Resolve(msg.SrcPathName)
		if srcPathName == "" {
			err := fmt.Errorf("invalid path name: %s", msg.SrcPathName)
			logger.Error(err)
			return true, err
		}
		src := filepath.Join(srcPathName, msg.Filename)
		// on any case, original file will be deleted
		defer func() {
			_ = os.Remove(src)
		}()

		dstPathName := config.Paths.Resolve(msg.DstPathName)
		if dstPathName == "" {
			err := fmt.Errorf("invalid path name: %s", msg.DstPathName)
			logger.Error(err)
			return true, err
		}
		dst := filepath.Join(dstPathName, msg.Filename) + ".m4a"

		srcAudio, err := ExtractAudio(src)
		if err != nil {
			logger.Error(err)
			return true, err
		}
		bitRate := srcAudio.BitRate
		minBitRate := msg.MinBitrate

		if !HasAudioEnoughQuality(bitRate, minBitRate) {
			err := fmt.Errorf("srcAudio bitdepth: %d is less than minimum: %d", bitRate, minBitRate)
			logger.Error(err)
			return true, err
		}

		// determine bitrate and convert srcAudio file to m4a
		bitRate = CalculateNumber(bitRate, msg.WantedBitRate)

		// convert file to desired format/resolution
		if err = commands.CmdFFMpegAudio(src, dst, bitRate); err != nil {
			logger.Error(err)
			return true, err
		}

		// extract media info as srcAudio from destination file
		dstAudio, err := ExtractAudio(dst)
		if err != nil {
			logger.Error(err)
			return true, err
		}
		info, err := os.Stat(dst)
		if err != nil {
			logger.Error(err)
			return true, err
		}
		dstAudio.TotalTrackCount = srcAudio.TotalTrackCount

		// write final srcAudio file to db
		file := &golang.File{
			ContentType: string(msg.DestinationContentType),
			Created:     time.Now(),
			Filename:    filepath.Base(dst),
			Id:          container.identifier.UUIDv4(),
			Length:      int(info.Size()),
			Status:      "converted",
		}
		audioFile := domains.NewAudioFile(dstAudio, file)
		err = container.repository.InsertAudioFile(&audioFile)
		return true, err
	}
}

func HasAudioEnoughQuality(bitRate, minBitrate int) bool {
	return bitRate >= minBitrate
}

// CalculateNumber returns the best match for a quality setting.
// If current is below setting, that value is used.
func CalculateNumber(current, wanted int) int {
	if current < wanted {
		return current
	}
	return wanted
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
