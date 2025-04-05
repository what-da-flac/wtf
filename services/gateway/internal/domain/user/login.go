package user

import (
	"context"
	"fmt"

	"github.com/what-da-flac/wtf/services/gateway/internal/helpers"
	"github.com/what-da-flac/wtf/services/gateway/internal/interfaces"

	"github.com/jinzhu/copier"
	"github.com/what-da-flac/wtf/openapi/gen/golang"
)

type Login struct {
	identifier interfaces.Identifier
	repository interfaces.Repository
	timer      interfaces.Timer
}

func NewLogin(identifier interfaces.Identifier, repository interfaces.Repository, timer interfaces.Timer) *Login {
	return &Login{
		repository: repository,
		identifier: identifier,
		timer:      timer,
	}
}

func (x *Login) validate(req *golang.PostV1UsersLoginJSONRequestBody) error {
	if x.repository == nil {
		return fmt.Errorf("missing parameter: repository")
	}
	if x.timer == nil {
		return fmt.Errorf("missing parameter: timer")
	}
	if x.identifier == nil {
		return fmt.Errorf("missing parameter: identifier")
	}
	u := req
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

func (x *Login) Login(ctx context.Context, req *golang.PostV1UsersLoginJSONRequestBody) (*golang.UserLoginResponse, error) {
	// validate incoming payload
	if err := x.validate(req); err != nil {
		return nil, err
	}
	res := &golang.UserLoginResponse{}
	now := x.timer.Now()
	// check if user already exists
	if user, err := x.repository.SelectUser(ctx, &req.Id, &req.Email); err == nil {
		user.Name = req.Name
		user.Image = req.Image
		user.LastLogin = x.timer.Now()
		if err := x.repository.UpdateUser(ctx, user); err != nil {
			return nil, err
		}
		if err = copier.Copy(res, user); err != nil {
			return nil, err
		}
		roleNames, err := helpers.ListRoleNamesForUser(ctx, x.repository, user)
		if err != nil {
			return nil, err
		}
		res.Roles = roleNames
		return res, nil
	}
	// prepare new user
	user := req
	user.Created = now
	user.LastLogin = now
	if err := x.repository.InsertUser(ctx, user); err != nil {
		return nil, err
	}
	if err := copier.Copy(res, user); err != nil {
		return nil, err
	}
	return res, nil
}
