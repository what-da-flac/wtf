package user

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

func (x *List) validate(req *golang.UserListParams) error {
	if x.repository == nil {
		return fmt.Errorf("missing parameter: repository")
	}
	if req == nil {
		return fmt.Errorf("missing parameter: req")
	}
	return nil
}

func (x *List) List(ctx context.Context, req *golang.UserListParams) ([]*golang.User, error) {
	if err := x.validate(req); err != nil {
		return nil, err
	}
	return x.repository.ListUsers(ctx, req)
}
