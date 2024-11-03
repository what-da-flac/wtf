package role

import (
	"context"
	"fmt"

	"github.com/what-da-flac/wtf/services/gateway/internal/interfaces"

	"github.com/what-da-flac/wtf/openapi/models"
)

type Update struct {
	repository interfaces.Repository
}

func NewUpdate(repository interfaces.Repository) *Update {
	return &Update{
		repository: repository,
	}
}

func (x *Update) validate(id string, role *models.RolePut) error {
	if x.repository == nil {
		return fmt.Errorf("missing parameter: repository")
	}
	if id == "" {
		return fmt.Errorf("missing parameter: id")
	}
	if role.Name == "" {
		return fmt.Errorf("missing parameter: name")
	}
	return nil
}

func (x *Update) Update(ctx context.Context, id string, role *models.RolePut) error {
	if err := x.validate(id, role); err != nil {
		return err
	}
	old, err := x.repository.SelectRole(ctx, id)
	if err != nil {
		return err
	}
	old.Description = role.Description
	old.Name = role.Name
	return x.repository.UpdateRole(ctx, old)
}
