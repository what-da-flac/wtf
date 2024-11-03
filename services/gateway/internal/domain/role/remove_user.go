package role

import (
	"context"
	"fmt"

	"github.com/what-da-flac/wtf/services/gateway/internal/interfaces"

	"github.com/what-da-flac/wtf/openapi/models"
)

type RemoveUser struct {
	repository interfaces.Repository
}

func NewRemoveUser(repository interfaces.Repository) *RemoveUser {
	return &RemoveUser{
		repository: repository,
	}
}

func (x *RemoveUser) validate(role *models.Role, user *models.User) error {
	if x.repository == nil {
		return fmt.Errorf("missing parameter: repository")
	}
	if role == nil || role.Id == "" {
		return fmt.Errorf("missing parameter: role")
	}
	if user == nil || user.Id == "" {
		return fmt.Errorf("missing parameter: user")
	}
	return nil
}

func (x *RemoveUser) Remove(ctx context.Context, role *models.Role, user *models.User) error {
	role, err := x.repository.SelectRole(ctx, role.Id)
	if err != nil {
		return fmt.Errorf("role %w", err)
	}
	user, err = x.repository.SelectUser(ctx, &user.Id, nil)
	if err != nil {
		return fmt.Errorf("user %w", err)
	}
	if err := x.validate(role, user); err != nil {
		return err
	}
	return x.repository.RemoveUser(ctx, role, user)
}
