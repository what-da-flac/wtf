package interfaces

import (
	"github.com/what-da-flac/wtf/openapi/models"
	"golang.org/x/net/context"
)

//go:generate moq -out ../../mocks/repository.go -pkg mocks . Repository
type Repository interface {
	InsertTorrent(ctx context.Context, t *models.Torrent) error
	InsertTorrentFile(ctx context.Context, t *models.TorrentFile) error
	DeleteTorrentFiles(ctx context.Context, id string) error
	SelectTorrent(ctx context.Context, id string) (*models.Torrent, error)
	SelectTorrentFiles(ctx context.Context, id string) ([]*models.TorrentFile, error)
	UpdateTorrent(ctx context.Context, t *models.Torrent) error
}
