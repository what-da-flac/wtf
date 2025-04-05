package helpers

import (
	"context"
	"fmt"

	"github.com/what-da-flac/wtf/openapi/gen/golang"

	"github.com/what-da-flac/wtf/services/gateway/internal/interfaces"
)

func ListRolesForUser(ctx context.Context, repository interfaces.Repository, user *golang.User) ([]*golang.Role, error) {
	var res []*golang.Role
	userRoles, err := repository.ListRolesForUser(ctx, user)
	if err != nil {
		return nil, err
	}
	roles, err := repository.ListRoles(ctx)
	if err != nil {
		return nil, err
	}
	keys := make(map[string]*golang.Role)
	for _, role := range roles {
		keys[role.Id] = role
	}
	for _, v := range userRoles {
		val, ok := keys[v.Id]
		if !ok {
			return nil, fmt.Errorf("cannot find role with id: %s", v.Id)
		}
		res = append(res, val)
	}
	return res, nil
}

func ListRoleNamesForUser(ctx context.Context, repository interfaces.Repository, user *golang.User) ([]string, error) {
	roles, err := ListRolesForUser(ctx, repository, user)
	if err != nil {
		return nil, err
	}
	return golang.Roles(roles).Names(), nil
}
