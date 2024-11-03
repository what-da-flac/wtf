package helpers

import (
	"context"

	"github.com/what-da-flac/wtf/services/gateway/internal/interfaces"

	"github.com/what-da-flac/wtf/openapi/models"
)

func UsersToMap(ctx context.Context, repository interfaces.Repository, ids ...string) (map[string]*models.User, error) {
	params := &models.UserListParams{
		Ids:         ids,
		OnlyDeleted: false,
	}
	rows, err := repository.ListUsers(ctx, params)
	if err != nil {
		return nil, err
	}
	return models.Users(rows).ToMap(), nil
}
