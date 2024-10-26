package user

import (
	"context"
	"fmt"

	"github.com/what-da-flac/wtf/gateway/internal/interfaces"
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

func (x *Delete) validate(id string) error {
	if id == "" {
		return fmt.Errorf("missing parameter: id")
	}
	return nil
}

func (x *Delete) Delete(ctx context.Context, id string) error {
	// validate incoming payload
	if err := x.validate(id); err != nil {
		return err
	}
	entity := NewLoad(x.repository)
	user, err := entity.Load(ctx, &models.User{Id: id})
	if err != nil {
		return err
	}
	user.IsDeleted = true
	return x.repository.UpdateUser(ctx, user)
}
