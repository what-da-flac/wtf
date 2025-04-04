package main

import (
	"encoding/json"
	"time"

	"github.com/what-da-flac/wtf/go-common/amazon"
	"github.com/what-da-flac/wtf/go-common/env"
	"github.com/what-da-flac/wtf/go-common/ifaces"
	"github.com/what-da-flac/wtf/go-common/loggers"
	"github.com/what-da-flac/wtf/go-common/rabbits"
	"github.com/what-da-flac/wtf/openapi/gen/golang"
	"github.com/what-da-flac/wtf/services/magnet-parser/internal/processors"
)

var Version string

func main() {
	logger := loggers.MustNewDevelopmentLogger()
	logger.Info("starting magnet-parser version:", Version)
	config := env.New()
	logger.Infof("trying to connect to rabbitmq at url: %s", config.RabbitMQ.URL)
	l := rabbits.NewListener(logger, env.QueueMagnetParser, config.RabbitMQ.URL, time.Second)
	defer func() { _ = l.Close() }()
	publisher := rabbits.NewPublisher(logger, env.QueueTorrentParser, config.RabbitMQ.URL)
	if err := publisher.Build(); err != nil {
		logger.Fatal(err)
	}
	defer func() { _ = publisher.Close() }()
	fn, err := processMessage(publisher, logger)
	if err != nil {
		logger.Fatal(err)
	}
	l.ListenAsync(fn)
	select {}
}

func processMessage(publisher ifaces.Publisher, logger ifaces.Logger) (func(msg []byte) (ack ifaces.AckType, err error), error) {
	awsSession := amazon.NewAWSSessionFromEnvironment()
	if err := awsSession.Build(); err != nil {
		return nil, err
	}
	cfg := awsSession.Session()
	return func(msg []byte) (ack ifaces.AckType, err error) {
		torrent := &golang.Torrent{}
		if err := json.Unmarshal(msg, torrent); err != nil {
			logger.Errorf("deserializing payload error: %v", err)
			return ifaces.MessageReject, nil
		}
		logger.Infof("received torrent with magnet link: %s", torrent.MagnetLink)
		if err := processors.Process(publisher, logger, cfg, torrent); err != nil {
			logger.Errorf("processing torrent error: %v", err)
			return ifaces.MessageReject, nil
		}
		return ifaces.MessageAcknowledge, nil
	}, nil
}
