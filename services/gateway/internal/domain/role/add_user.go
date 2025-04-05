package role

import (
	"context"
	"fmt"
	"strings"

	"github.com/what-da-flac/wtf/services/gateway/internal/interfaces"

	"github.com/what-da-flac/wtf/openapi/gen/golang"
)

type AddUser struct {
	repository interfaces.Repository
}

func NewAddUser(repository interfaces.Repository) *AddUser {
	return &AddUser{
		repository: repository,
	}
}

func (x *AddUser) validate(role *golang.Role, user *golang.User) error {
	if role == nil || role.Id == "" {
		return fmt.Errorf("missing parameter: role")
	}
	if user == nil || user.Id == "" {
		return fmt.Errorf("missing parameter: user")
	}
	return nil
}

func (x *AddUser) Add(ctx context.Context, role *golang.Role, user *golang.User) error {
	const constraintText = "constraint"
	if err := x.validate(role, user); err != nil {
		return err
	}
	role, err := x.repository.SelectRole(ctx, role.Id)
	if err != nil {
		return fmt.Errorf("role %w", err)
	}
	user, err = x.repository.SelectUser(ctx, &user.Id, nil)
	if err != nil {
		return fmt.Errorf("user %w", err)
	}
	if err = x.repository.AddUser(ctx, role, user); err != nil {
		if strings.Contains(err.Error(), constraintText) {
			return err
		}
	}
	return nil
}
