package torrent

import (
	"context"

	interfaces2 "github.com/what-da-flac/wtf/services/gateway/internal/interfaces"

	"github.com/what-da-flac/wtf/openapi/models"
)

type Update struct {
	identifier interfaces2.Identifier
	repository interfaces2.Repository
	timer      interfaces2.Timer
}

func NewUpdate(repository interfaces2.Repository, timer interfaces2.Timer,
	identifier interfaces2.Identifier) *Update {
	return &Update{
		identifier: identifier,
		repository: repository,
		timer:      timer,
	}
}

func (x *Update) Save(ctx context.Context, torrent *models.Torrent) error {
	old, err := x.repository.SelectTorrent(ctx, torrent.Id)
	if err != nil {
		return err
	}
	torrent.Created = old.Created
	if len(torrent.Files) != 0 {
		// delete previous files
		if err := x.repository.DeleteTorrentFiles(ctx, torrent.Id); err != nil {
			return err
		}
		// insert new torrent files
		for _, file := range torrent.Files {
			file.Id = x.identifier.UUIDv4()
			file.TorrentId = torrent.Id
			if err := x.repository.InsertTorrentFile(ctx, &file); err != nil {
				return err
			}
		}
	}
	now := x.timer.Now()
	torrent.Updated = &now
	return x.repository.UpdateTorrent(ctx, torrent)
}
