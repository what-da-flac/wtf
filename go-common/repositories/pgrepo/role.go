package pgrepo

import (
	"context"

	"github.com/what-da-flac/wtf/openapi/gen/golang"
)

func (x *PgRepo) InsertRole(_ context.Context, role *golang.Role) error {
	db := x.GORM()
	return db.Create(roleFromProto(role)).Error
}

func (x *PgRepo) UpdateRole(_ context.Context, role *golang.Role) error {
	db := x.GORM()
	return db.Updates(roleFromProto(role)).Error
}

func (x *PgRepo) SelectRole(_ context.Context, id string) (*golang.Role, error) {
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

func (x *PgRepo) ListRoles(_ context.Context) ([]*golang.Role, error) {
	var (
		result []*golang.Role
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
