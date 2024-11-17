package main

import (
	"encoding/json"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/what-da-flac/wtf/go-common/env"
	"github.com/what-da-flac/wtf/go-common/identifiers"
	"github.com/what-da-flac/wtf/go-common/ifaces"
	"github.com/what-da-flac/wtf/go-common/loggers"
	"github.com/what-da-flac/wtf/go-common/pgpq"
	"github.com/what-da-flac/wtf/go-common/rabbits"
	"github.com/what-da-flac/wtf/go-common/repositories/pgrepo"
	"github.com/what-da-flac/wtf/openapi/models"
	"github.com/what-da-flac/wtf/services/torrent-info/internal/interfaces"
	"golang.org/x/net/context"
)

var Version string

func main() {
	logger := loggers.MustNewDevelopmentLogger()
	logger.Info("starting torrent-info:", Version)
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
	identifier := identifiers.NewIdentifier()
	fn, err := processMessage(logger, repo, identifier)
	if err != nil {
		logger.Error(err)
	}
	l.ListenAsync(fn)
	select {}
}

func processMessage(logger ifaces.Logger, repo interfaces.Repository,
	identifier ifaces.Identifier,
) (func(msg []byte) (ack ifaces.AckType, err error), error) {
	return func(msg []byte) (ack ifaces.AckType, err error) {
		ctx := context.TODO()
		torrent := &models.Torrent{}
		if err := json.Unmarshal(msg, torrent); err != nil {
			logger.Errorf("deserializing payload error: %v", err)
			return ifaces.MessageReject, nil
		}
		logger.Infof("received torrent with filename: %s", torrent.Filename)
		if err = upsertTorrent(ctx, repo, identifier, torrent); err != nil {
			logger.Errorf("processing torrent error: %v", err)
			return ifaces.MessageReject, nil
		}
		return ifaces.MessageAcknowledge, nil
	}, nil
}

func upsertTorrent(ctx context.Context, repo interfaces.Repository, identifier ifaces.Identifier, t *models.Torrent) error {
	if t == nil || t.Id == "" {
		return fmt.Errorf("missing id in torrent, cannot process in db")
	}
	if t.User == nil || t.User.Id == "" {
		return fmt.Errorf("missing user in torrent, cannot process in db")
	}
	if _, err := repo.SelectTorrent(ctx, t.Id); err != nil {
		if err = repo.InsertTorrent(ctx, t); err != nil {
			return err
		}
	} else {
		if err = repo.UpdateTorrent(ctx, t); err != nil {
			return err
		}
	}
	// check if there are no files for this torrent, insert
	files, err := repo.SelectTorrentFiles(ctx, t.Id)
	if err != nil {
		return err
	}
	if len(files) != 0 {
		return nil
	}
	// delete previous files
	if err = repo.DeleteTorrentFiles(ctx, t.Id); err != nil {
		return err
	}
	// insert current files
	for _, file := range t.Files {
		file.Id = identifier.UUIDv4()
		file.TorrentId = t.Id
		if err = repo.InsertTorrentFile(ctx, &file); err != nil {
			return err
		}
	}
	return nil
}
