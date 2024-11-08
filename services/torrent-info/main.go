package main

import (
	"encoding/json"
	"time"

	"github.com/what-da-flac/wtf/go-common/env"
	"github.com/what-da-flac/wtf/go-common/ifaces"
	"github.com/what-da-flac/wtf/go-common/loggers"
	"github.com/what-da-flac/wtf/go-common/rabbits"
	"github.com/what-da-flac/wtf/openapi/models"
	"github.com/what-da-flac/wtf/services/torrent-info/internal/processors"
)

var (
	Version string
)

func main() {
	logger := loggers.MustNewDevelopmentLogger()
	logger.Info("starting torrent-parser version:", Version)
	config := env.New()
	l := rabbits.NewListener(logger, env.QueueTorrentInfo, config.RabbitMQ.URL, time.Second)
	defer func() { _ = l.Close() }()
	fn, err := processMessage(logger, config)
	if err != nil {
		logger.Error(err)
	}
	l.ListenAsync(fn)
	select {}
}

func processMessage(logger ifaces.Logger, config *env.Config) (func(msg []byte) (ack ifaces.AckType, err error), error) {
	return func(msg []byte) (ack ifaces.AckType, err error) {
		torrent := &models.Torrent{}
		if err := json.Unmarshal(msg, torrent); err != nil {
			logger.Errorf("deserializing payload error: %v", err)
			return ifaces.MessageReject, nil
		}
		logger.Infof("received torrent with filename: %s", torrent.Filename)
		if err := processors.Process(logger, config, torrent); err != nil {
			logger.Errorf("processing torrent error: %v", err)
			logger.Info("sending to queue again until this is fully implemented ")
			// TODO: reject or requeue, but don't ignore this
			return ifaces.MessageAcknowledge, nil
		}
		return ifaces.MessageAcknowledge, nil
	}, nil
}
