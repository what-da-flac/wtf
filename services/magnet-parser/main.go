package main

import (
	"time"

	"github.com/what-da-flac/wtf/go-common/environment"
	"github.com/what-da-flac/wtf/go-common/ifaces"
	"github.com/what-da-flac/wtf/go-common/loggers"
	"github.com/what-da-flac/wtf/go-common/rabbits"
)

var (
	Version string
)

func main() {
	logger := loggers.MustNewDevelopmentLogger()
	logger.Info("starting magnet-parser")
	config := environment.New()
	l := rabbits.NewListener(logger, "test-queue", config.RabbitMQ.URL, time.Second)
	defer func() { _ = l.Close() }()
	l.ListenAsync(func(msg []byte) (ack ifaces.AckType, err error) {
		logger.Info("received:", string(msg))
		return ifaces.MessageAcknowledge, nil
	})
	select {}
}
