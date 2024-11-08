package pgrepo

import (
	"github.com/jinzhu/copier"
	"github.com/what-da-flac/wtf/openapi/models"
)

type RoleDto struct {
	Id          string
	Name        string
	Description string
}

func (x *RoleDto) TableName() string { return "roles" }

func (x *RoleDto) toProto() *models.Role {
	result := &models.Role{}
	_ = copier.Copy(result, x)
	return result
}

func roleFromProto(role *models.Role) *RoleDto {
	res := &RoleDto{}
	if err := copier.Copy(res, role); err != nil {
		return nil
	}
	return res
}
