package main

import (
	"encoding/json"
	"time"

	"github.com/what-da-flac/wtf/go-common/pgpq"
	"github.com/what-da-flac/wtf/go-common/repositories/pgrepo"

	"golang.org/x/net/context"

	"github.com/what-da-flac/wtf/go-common/env"
	"github.com/what-da-flac/wtf/go-common/ifaces"
	"github.com/what-da-flac/wtf/go-common/loggers"
	"github.com/what-da-flac/wtf/go-common/rabbits"
	"github.com/what-da-flac/wtf/openapi/models"
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
	dbUri := config.DB.URL
	db, err := pgpq.New(dbUri)
	if err != nil {
		panic(err)
	}
	repo, err := pgrepo.NewPgRepo(db, dbUri, false)
	if err != nil {
		panic(err)
	}
	fn, err := processMessage(logger, repo)
	if err != nil {
		logger.Error(err)
	}
	l.ListenAsync(fn)
	select {}
}

func processMessage(logger ifaces.Logger, repo interface {
	InsertTorrent(context.Context, *models.Torrent) error
}) (func(msg []byte) (ack ifaces.AckType, err error), error) {
	return func(msg []byte) (ack ifaces.AckType, err error) {
		torrent := &models.Torrent{}
		if err := json.Unmarshal(msg, torrent); err != nil {
			logger.Errorf("deserializing payload error: %v", err)
			return ifaces.MessageReject, nil
		}
		logger.Infof("received torrent with filename: %s", torrent.Filename)
		if err = repo.InsertTorrent(context.Background(), torrent); err != nil {
			logger.Errorf("processing torrent error: %v", err)
			return ifaces.MessageReject, nil
		}
		return ifaces.MessageAcknowledge, nil
	}, nil
}
