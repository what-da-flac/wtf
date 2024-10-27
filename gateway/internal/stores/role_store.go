package stores

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/what-da-flac/wtf/gateway/internal/helpers"
	"github.com/what-da-flac/wtf/gateway/internal/interfaces"
	"github.com/what-da-flac/wtf/go-common/imodel"
	"github.com/what-da-flac/wtf/openapi/models"
)

type RoleStore struct {
	repository interfaces.Repository
}

func NewRoleStore(repository interfaces.Repository) *RoleStore {
	return &RoleStore{
		repository: repository,
	}
}

func (x *RoleStore) Roles(ctx context.Context, user *imodel.User) ([]string, error) {
	u := &models.User{}
	if err := copier.Copy(u, user); err != nil {
		return nil, err
	}
	return helpers.ListRoleNamesForUser(ctx, x.repository, u)
}