package torrent

import (
	"context"
	"fmt"

	"github.com/what-da-flac/wtf/gateway/internal/helpers"

	"github.com/what-da-flac/wtf/gateway/internal/interfaces"
	"github.com/what-da-flac/wtf/openapi/models"
)

type List struct {
	repository interfaces.Repository
}

func NewList(repository interfaces.Repository) *List {
	return &List{
		repository: repository,
	}
}

func (x *List) List(ctx context.Context, params models.GetV1TorrentsParams) ([]*models.Torrent, error) {
	var ids []string
	rows, err := x.repository.ListTorrents(ctx, params)
	if err != nil {
		return nil, err
	}
	for _, row := range rows {
		ids = append(ids, row.User.Id)
	}
	users, err := helpers.UsersToMap(ctx, x.repository, ids...)
	if err != nil {
		return nil, err
	}
	for _, row := range rows {
		id := row.User.Id
		user, exists := users[id]
		if !exists {
			return nil, fmt.Errorf("user not found with id: %s", id)
		}
		row.User = user
	}
	return rows, nil
}
