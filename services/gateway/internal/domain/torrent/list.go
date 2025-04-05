package torrent

import (
	"context"
	"fmt"

	"github.com/what-da-flac/wtf/services/gateway/internal/helpers"
	"github.com/what-da-flac/wtf/services/gateway/internal/interfaces"

	"github.com/what-da-flac/wtf/openapi/gen/golang"
)

type List struct {
	repository interfaces.Repository
}

func NewList(repository interfaces.Repository) *List {
	return &List{
		repository: repository,
	}
}

func (*List) validate(params golang.GetV1TorrentsParams) error {
	if params.Limit < 1 {
		return fmt.Errorf("limit must be greater than 0")
	}
	if params.SortField == "" {
		return fmt.Errorf("sort field must be set")
	}
	if params.SortDirection != "asc" && params.SortDirection != "desc" {
		return fmt.Errorf("sort direction must be asc or desc")
	}
	return nil
}

func (x *List) List(ctx context.Context, params golang.GetV1TorrentsParams) ([]*golang.Torrent, error) {
	var ids []string
	if err := x.validate(params); err != nil {
		return nil, err
	}
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
