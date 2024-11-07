package main

import (
	"encoding/json"
	"time"

	"github.com/what-da-flac/wtf/go-common/environment"
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
	config := environment.New()
	l := rabbits.NewListener(logger, config.Queues.TorrentInfo, config.RabbitMQ.URL, time.Second)
	defer func() { _ = l.Close() }()
	fn, err := processMessage(logger, config)
	if err != nil {
		logger.Fatal(err)
	}
	l.ListenAsync(fn)
	select {}
}

// TODO: figure out this error when running all queues

/*
torrent-info    | 2024-11-07T15:35:27.748Z	INFO	torrent-info/main.go:40	received torrent with filename: ded35714499d53f95d5b382e522abcfbd7c67144.torrent
torrent-info    | 2024-11-07T15:35:27.748Z	INFO	processors/process.go:10	torrent name: Coldplay - Moon Music (2024) FLAC [PMEDIA] ⭐️ file count: 12
torrent-info    | 2024-11-07T15:35:27.748Z	WARN	processors/process.go:11	TODO: store in database
torrent-info    | github.com/what-da-flac/wtf/services/torrent-info/internal/processors.Process
torrent-info    | 	/app/services/torrent-info/internal/processors/process.go:11
torrent-info    | main.main.processMessage.func2
torrent-info    | 	/app/services/torrent-info/main.go:41
torrent-info    | github.com/what-da-flac/wtf/go-common/rabbits.(*Listener).processMessage
torrent-info    | 	/app/go-common/rabbits/listener.go:124
torrent-info    | github.com/what-da-flac/wtf/go-common/rabbits.(*Listener).listen
torrent-info    | 	/app/go-common/rabbits/listener.go:99
torrent-info    | github.com/what-da-flac/wtf/go-common/rabbits.(*Listener).ListenAsync.func1
torrent-info    | 	/app/go-common/rabbits/listener.go:60
*/

func processMessage(logger ifaces.Logger, config *environment.Config) (func(msg []byte) (ack ifaces.AckType, err error), error) {
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
			return ifaces.MessageReject, nil
		}
		return ifaces.MessageAcknowledge, nil
	}, nil
}
