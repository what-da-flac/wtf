package user

import (
	"context"
	"fmt"

	interfaces2 "github.com/what-da-flac/wtf/services/gateway/internal/interfaces"

	"github.com/what-da-flac/wtf/openapi/models"
)

type Create struct {
	identifier interfaces2.Identifier
	repository interfaces2.Repository
	timer      interfaces2.Timer
}

func NewCreate(identifier interfaces2.Identifier, repository interfaces2.Repository, timer interfaces2.Timer) *Create {
	return &Create{
		identifier: identifier,
		repository: repository,
	}
}

func (x *Create) validate(u *models.UserPost) error {
	if x.identifier == nil {
		return fmt.Errorf("missing parameter: identifier")
	}
	if x.timer == nil {
		return fmt.Errorf("missing parameter: timer")
	}
	if x.repository == nil {
		return fmt.Errorf("missing parameter: repository")
	}
	if value := u.Id; value == "" {
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

func (x *Create) Save(ctx context.Context, req *models.UserPost) (*models.User, error) {
	// validate incoming payload
	if err := x.validate(req); err != nil {
		return nil, err
	}
	// check if user already exists
	if _, err := x.repository.SelectUser(ctx, nil, &req.Email); err == nil {
		return nil, fmt.Errorf("user is already registered with email: %s", req.Email)
	}
	// insert new user
	now := x.timer.Now()
	user := &models.User{
		Id:        req.Id,
		Created:   now,
		LastLogin: now,
	}
	if user.Id == "" {
		user.Id = x.identifier.UUIDv4()
	}
	if err := x.repository.InsertUser(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}
