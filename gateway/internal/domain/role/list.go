package role

import (
	"context"
	"fmt"

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

func (x *List) validate() error {
	if x.repository == nil {
		return fmt.Errorf("missing parameter: repository")
	}
	return nil
}

func (x *List) List(ctx context.Context) ([]*models.Role, error) {
	if err := x.validate(); err != nil {
		return nil, err
	}
	return x.repository.ListRoles(ctx)
}
