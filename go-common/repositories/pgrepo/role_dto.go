package pgrepo

import (
	"github.com/jinzhu/copier"
	"github.com/what-da-flac/wtf/openapi/gen/golang"
)

type RoleDto struct {
	Id          string
	Name        string
	Description string
}

func (x *RoleDto) TableName() string { return "roles" }

func (x *RoleDto) toProto() *golang.Role {
	result := &golang.Role{}
	_ = copier.Copy(result, x)
	return result
}

func roleFromProto(role *golang.Role) *RoleDto {
	res := &RoleDto{}
	if err := copier.Copy(res, role); err != nil {
		return nil
	}
	return res
}
