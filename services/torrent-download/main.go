package main

import (
	"encoding/json"
	"os"
	"time"

	"github.com/what-da-flac/wtf/go-common/amazon"
	"github.com/what-da-flac/wtf/go-common/env"
	"github.com/what-da-flac/wtf/go-common/ifaces"
	"github.com/what-da-flac/wtf/go-common/loggers"
	"github.com/what-da-flac/wtf/go-common/rabbits"
	"github.com/what-da-flac/wtf/openapi/models"
	"github.com/what-da-flac/wtf/services/torrent-download/internal/downloaders"
	"github.com/what-da-flac/wtf/services/torrent-download/internal/processors"
)

var Version string

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	logger := loggers.MustNewDevelopmentLogger()
	logger.Info("starting torrent-parser version:", Version)
	config := env.New()
	logger.Infof("trying to connect to rabbitmq at url: %s", config.RabbitMQ.URL)
	l := rabbits.NewListener(logger, env.QueueTorrentDownload, config.RabbitMQ.URL, time.Second)
	defer func() { _ = l.Close() }()
	publisher := rabbits.NewPublisher(logger, env.QueueTorrentInfo, config.RabbitMQ.URL)
	if err := publisher.Build(); err != nil {
		logger.Fatal(err)
	}
	defer func() { _ = publisher.Close() }()
	fn, err := processMessage(logger, config, publisher)
	if err != nil {
		logger.Fatal(err)
	}
	l.ListenAsync(fn)
	select {}

}

func processMessage(logger ifaces.Logger, config *env.Config, publisher ifaces.Publisher) (func(msg []byte) (ack ifaces.AckType, err error), error) {
	awsSession := amazon.NewAWSSessionFromEnvironment()
	if err := awsSession.Build(); err != nil {
		return nil, err
	}
	sess := awsSession.Session()
	s3Downloader := amazon.NewDownloader(sess)
	torrentDownloader := downloaders.NewTorrentDownloader(logger, config.Downloads.Timeout)
	return func(msg []byte) (ack ifaces.AckType, err error) {
		torrent := &models.Torrent{}
		if err := json.Unmarshal(msg, torrent); err != nil {
			logger.Errorf("deserializing payload error: %v", err)
			return ifaces.MessageReject, nil
		}
		logger.Infof("received torrent with filename: %s", torrent.Filename)
		elapsed, err := processors.Process(logger, torrentDownloader, s3Downloader, torrent, config, os.TempDir())
		if err != nil {
			logger.Errorf("processing torrent error: %v", err)
			return ifaces.MessageReject, nil
		}
		logger.Infof("downloaded torrent, time elapsed: %v", elapsed)
		torrent.Status = models.Saved
		data, err := json.Marshal(torrent)
		if err != nil {
			logger.Errorf("marshaling torrent error: %v", err)
			return ifaces.MessageReject, nil
		}
		if err = publisher.Publish(data); err != nil {
			logger.Errorf("publishing torrent info error: %v", err)
			return ifaces.MessageReject, nil
		}
		return ifaces.MessageAcknowledge, nil
	}, nil
}
