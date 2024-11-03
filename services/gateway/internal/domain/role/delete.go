package role

import (
	"context"
	"fmt"

	"github.com/what-da-flac/wtf/services/gateway/internal/interfaces"

	"github.com/what-da-flac/wtf/openapi/models"
)

type Delete struct {
	repository interfaces.Repository
}

func NewDelete(repository interfaces.Repository) *Delete {
	return &Delete{
		repository: repository,
	}
}

func (x *Delete) validate(role *models.Role) error {
	if x.repository == nil {
		return fmt.Errorf("missing parameter: repository")
	}
	if role == nil || role.Id == "" {
		return fmt.Errorf("missing parameter: id")
	}
	return nil
}

func (x *Delete) Delete(ctx context.Context, role *models.Role) error {
	if err := x.validate(role); err != nil {
		return err
	}
	role, err := x.repository.SelectRole(ctx, role.Id)
	if err != nil {
		return err
	}
	rows, err := x.repository.ListUsersInRole(ctx, role.Id)
	if err != nil {
		return err
	}
	if len(rows) != 0 {
		return fmt.Errorf("cannot delete role: %s there are related users", role.Name)
	}
	return x.repository.DeleteRole(ctx, role.Id)
}
