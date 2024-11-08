package pgrepo

import (
	"context"

	"github.com/what-da-flac/wtf/openapi/models"
)

func (x *PgRepo) InsertRole(_ context.Context, role *models.Role) error {
	db := x.GORM()
	return db.Create(roleFromProto(role)).Error
}

func (x *PgRepo) UpdateRole(_ context.Context, role *models.Role) error {
	db := x.GORM()
	return db.Updates(roleFromProto(role)).Error
}

func (x *PgRepo) SelectRole(_ context.Context, id string) (*models.Role, error) {
	db := x.GORM()
	row := &RoleDto{
		Id: id,
	}
	if err := db.First(row).Error; err != nil {
		return nil, err
	}
	return row.toProto(), nil
}

func (x *PgRepo) DeleteRole(_ context.Context, id string) error {
	db := x.GORM()
	row := &RoleDto{
		Id: id,
	}
	return db.Delete(row).Unscoped().Error
}

func (x *PgRepo) ListRoles(_ context.Context) ([]*models.Role, error) {
	var (
		result []*models.Role
		rows   []RoleDto
	)
	db := x.GORM()
	if err := db.Find(&rows).Error; err != nil {
		return nil, err
	}
	for _, row := range rows {
		result = append(result, row.toProto())
	}
	return result, nil
}
