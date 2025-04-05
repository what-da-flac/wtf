package role

import (
	"context"
	"fmt"

	"github.com/what-da-flac/wtf/services/gateway/internal/interfaces"

	"github.com/what-da-flac/wtf/openapi/gen/golang"
)

type ListUsers struct {
	repository interfaces.Repository
}

func NewListUsers(repository interfaces.Repository) *ListUsers {
	return &ListUsers{
		repository: repository,
	}
}

func (x *ListUsers) validate(roleId string) error {
	if x.repository == nil {
		return fmt.Errorf("missing parameter: repository")
	}
	if roleId == "" {
		return fmt.Errorf("missing parameter: role")
	}
	return nil
}

func (x *ListUsers) List(ctx context.Context, roleId string) ([]*golang.User, error) {
	var (
		ids    []string
		result []*golang.User
	)
	if err := x.validate(roleId); err != nil {
		return nil, err
	}
	roleUsers, err := x.repository.ListUsersInRole(ctx, roleId)
	if err != nil {
		return nil, err
	}
	for _, v := range roleUsers {
		ids = append(ids, v.Id)
	}
	users, err := x.repository.ListUsers(ctx, &golang.UserListParams{
		Ids: ids,
	})
	if err != nil {
		return nil, err
	}
	keys := make(map[string]*golang.User)
	for _, user := range users {
		keys[user.Id] = user
	}
	for _, v := range roleUsers {
		val, ok := keys[v.Id]
		if !ok {
			return nil, fmt.Errorf("cannot find user with id: %s", v.Id)
		}
		result = append(result, val)
	}
	return result, nil
}
