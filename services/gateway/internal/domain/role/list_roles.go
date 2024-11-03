package role

import (
	"context"
	"fmt"

	"github.com/what-da-flac/wtf/services/gateway/internal/helpers"
	"github.com/what-da-flac/wtf/services/gateway/internal/interfaces"

	"github.com/what-da-flac/wtf/openapi/models"
)

type ListRoles struct {
	repository interfaces.Repository
}

func NewListRoles(repository interfaces.Repository) *ListRoles {
	return &ListRoles{
		repository: repository,
	}
}

func (x *ListRoles) validate(user *models.User) error {
	if x.repository == nil {
		return fmt.Errorf("missing parameter: repository")
	}
	if user == nil || user.Id == "" {
		return fmt.Errorf("missing parameter: user")
	}
	return nil
}

func (x *ListRoles) List(ctx context.Context, user *models.User) ([]*models.Role, error) {
	if err := x.validate(user); err != nil {
		return nil, err
	}
	return helpers.ListRolesForUser(ctx, x.repository, user)
}
