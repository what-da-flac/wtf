package role

import (
	"context"
	"fmt"

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

func (x *List) validate() error {
	if x.repository == nil {
		return fmt.Errorf("missing parameter: repository")
	}
	return nil
}

func (x *List) List(ctx context.Context) ([]*golang.Role, error) {
	if err := x.validate(); err != nil {
		return nil, err
	}
	return x.repository.ListRoles(ctx)
}
