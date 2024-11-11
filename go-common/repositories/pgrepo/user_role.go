package pgrepo

import (
	"context"

	"github.com/what-da-flac/wtf/openapi/models"
)

func (x *PgRepo) AddUser(_ context.Context, role *models.Role, user *models.User) error {
	db := x.GORM()
	return db.Create(&UserRoleDto{
		RoleId: role.Id,
		UserId: user.Id,
	}).Error
}
func (x *PgRepo) RemoveUser(_ context.Context, role *models.Role, user *models.User) error {
	model := &UserRoleDto{
		RoleId: role.Id,
		UserId: user.Id,
	}
	db := x.GORM()
	return db.Where(model).Delete(&UserRoleDto{}).Unscoped().Error
}

func (x *PgRepo) ListUsersInRole(_ context.Context, roleId string) ([]*models.User, error) {
	var (
		result []*models.User
		rows   []*UserRoleDto
	)
	db := x.GORM()
	if err := db.Where(&UserRoleDto{RoleId: roleId}).Find(&rows).Error; err != nil {
		return nil, err
	}
	for _, row := range rows {
		result = append(result, &models.User{Id: row.UserId})
	}
	return result, nil
}

func (x *PgRepo) ListRolesForUser(_ context.Context, user *models.User) ([]*models.Role, error) {
	var (
		result []*models.Role
		rows   []*UserRoleDto
	)
	db := x.GORM()
	if err := db.Where(&UserRoleDto{UserId: user.Id}).Find(&rows).Error; err != nil {
		return nil, err
	}
	for _, row := range rows {
		result = append(result, &models.Role{Id: row.RoleId})
	}
	return result, nil
}
