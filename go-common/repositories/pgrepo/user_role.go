package pgrepo

import (
	"context"

	"github.com/what-da-flac/wtf/openapi/gen/golang"
)

func (x *PgRepo) AddUser(_ context.Context, role *golang.Role, user *golang.User) error {
	db := x.GORM()
	return db.Create(&UserRoleDto{
		RoleId: role.Id,
		UserId: user.Id,
	}).Error
}
func (x *PgRepo) RemoveUser(_ context.Context, role *golang.Role, user *golang.User) error {
	model := &UserRoleDto{
		RoleId: role.Id,
		UserId: user.Id,
	}
	db := x.GORM()
	return db.Where(model).Delete(&UserRoleDto{}).Unscoped().Error
}

func (x *PgRepo) ListUsersInRole(_ context.Context, roleId string) ([]*golang.User, error) {
	var (
		result []*golang.User
		rows   []*UserRoleDto
	)
	db := x.GORM()
	if err := db.Where(&UserRoleDto{RoleId: roleId}).Find(&rows).Error; err != nil {
		return nil, err
	}
	for _, row := range rows {
		result = append(result, &golang.User{Id: row.UserId})
	}
	return result, nil
}

func (x *PgRepo) ListRolesForUser(_ context.Context, user *golang.User) ([]*golang.Role, error) {
	var (
		result []*golang.Role
		rows   []*UserRoleDto
	)
	db := x.GORM()
	if err := db.Where(&UserRoleDto{UserId: user.Id}).Find(&rows).Error; err != nil {
		return nil, err
	}
	for _, row := range rows {
		result = append(result, &golang.Role{Id: row.RoleId})
	}
	return result, nil
}
