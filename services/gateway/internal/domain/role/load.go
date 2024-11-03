package role

import (
	"context"
	"fmt"

	"github.com/what-da-flac/wtf/services/gateway/internal/interfaces"

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

func (x *Load) validate(id string) error {
	if x.repository == nil {
		return fmt.Errorf("missing parameter: repository")
	}
	if id == "" {
		return fmt.Errorf("missing parameter: id")
	}
	return nil
}

func (x *Load) Load(ctx context.Context, id string) (*models.Role, error) {
	if err := x.validate(id); err != nil {
		return nil, err
	}
	return x.repository.SelectRole(ctx, id)
}
