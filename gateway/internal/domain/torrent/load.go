package torrent

import (
	"context"

	"github.com/what-da-flac/wtf/gateway/internal/interfaces"
	"github.com/what-da-flac/wtf/openapi/models"
)

type Load struct {
	repository interfaces.Repository
}

func NewLoad(repository interfaces.Repository) *Load {
	return &Load{
		repository: repository,
	}
}

func (x *Load) Load(ctx context.Context, id string) (*models.Torrent, error) {
	res, err := x.repository.SelectTorrent(ctx, id)
	if err != nil {
		return nil, err
	}
	user, err := x.repository.SelectUser(ctx, &res.User.Id, &res.User.Email)
	if err != nil {
		return nil, err
	}
	res.User = user
	files, err := x.repository.SelectTorrentFiles(ctx, res.Id)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		res.Files = append(res.Files, *file)
	}
	return res, nil
}
