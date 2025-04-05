package interfaces

import (
	"github.com/what-da-flac/wtf/openapi/gen/golang"
	"golang.org/x/net/context"
)

//go:generate moq -out ../../mocks/repository.go -pkg mocks . Repository
type Repository interface {
	InsertTorrent(ctx context.Context, t *golang.Torrent) error
	InsertTorrentFile(ctx context.Context, t *golang.TorrentFile) error
	DeleteTorrentFiles(ctx context.Context, id string) error
	SelectTorrent(ctx context.Context, id string) (*golang.Torrent, error)
	SelectTorrentFiles(ctx context.Context, id string) ([]*golang.TorrentFile, error)
	UpdateTorrent(ctx context.Context, t *golang.Torrent) error
}
