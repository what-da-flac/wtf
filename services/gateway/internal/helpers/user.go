package helpers

import (
	"context"

	"github.com/what-da-flac/wtf/openapi/gen/golang"

	"github.com/what-da-flac/wtf/services/gateway/internal/interfaces"
)

func UsersToMap(ctx context.Context, repository interfaces.Repository, ids ...string) (map[string]*golang.User, error) {
	params := &golang.UserListParams{
		Ids:         ids,
		OnlyDeleted: false,
	}
	rows, err := repository.ListUsers(ctx, params)
	if err != nil {
		return nil, err
	}
	return golang.Users(rows).ToMap(), nil
}
