package user

import (
	"context"
	"fmt"

	interfaces2 "github.com/what-da-flac/wtf/services/gateway/internal/interfaces"

	"github.com/what-da-flac/wtf/openapi/models"
)

type Update struct {
	repository interfaces2.Repository
	timer      interfaces2.Timer
}

func NewUpdate(repository interfaces2.Repository, timer interfaces2.Timer) *Update {
	return &Update{
		repository: repository,
		timer:      timer,
	}
}

func (x *Update) validate(id string, req *models.UserPut) error {
	if x.repository == nil {
		return fmt.Errorf("missing parameter: repository")
	}
	if x.timer == nil {
		return fmt.Errorf("missing parameter: timer")
	}
	u := req
	if id == "" {
		return fmt.Errorf("missing parameter: id")
	}
	if value := u.Name; value == "" {
		return fmt.Errorf("missing parameter: name")
	}
	if value := u.Email; value == "" {
		return fmt.Errorf("missing parameter: email")
	}
	return nil
}

func (x *Update) Save(ctx context.Context, id string, req *models.UserPut) (*models.User, error) {
	// validate incoming payload
	if err := x.validate(id, req); err != nil {
		return nil, err
	}
	user, err := NewLoad(x.repository).Load(ctx, &models.User{Id: id})
	if err != nil {
		return nil, err
	}
	user.LastLogin = x.timer.Now()
	user.Name = req.Name
	user.Image = req.Image
	if err = x.repository.UpdateUser(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}
