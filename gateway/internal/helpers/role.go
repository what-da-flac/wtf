package helpers

import (
	"context"
	"fmt"

	"github.com/what-da-flac/wtf/gateway/internal/interfaces"
	"github.com/what-da-flac/wtf/openapi/models"
)

func ListRolesForUser(ctx context.Context, repository interfaces.Repository, user *models.User) ([]*models.Role, error) {
	var res []*models.Role
	userRoles, err := repository.ListRolesForUser(ctx, user)
	if err != nil {
		return nil, err
	}
	roles, err := repository.ListRoles(ctx)
	if err != nil {
		return nil, err
	}
	keys := make(map[string]*models.Role)
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

func ListRoleNamesForUser(ctx context.Context, repository interfaces.Repository, user *models.User) ([]string, error) {
	roles, err := ListRolesForUser(ctx, repository, user)
	if err != nil {
		return nil, err
	}
	return models.Roles(roles).Names(), nil
}
