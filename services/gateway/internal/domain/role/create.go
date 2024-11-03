package role

import (
	"context"
	"fmt"
	"strings"

	interfaces2 "github.com/what-da-flac/wtf/services/gateway/internal/interfaces"

	"github.com/what-da-flac/wtf/openapi/models"
)

type Create struct {
	repository interfaces2.Repository
	identifier interfaces2.Identifier
}

func NewCreate(identifier interfaces2.Identifier, repository interfaces2.Repository) *Create {
	return &Create{
		identifier: identifier,
		repository: repository,
	}
}

func (x *Create) validate(role *models.PostV1RolesJSONRequestBody) error {
	if x.repository == nil {
		return fmt.Errorf("missing parameter: repository")
	}
	if x.identifier == nil {
		return fmt.Errorf("missing parameter: identifier")
	}
	if role == nil {
		return fmt.Errorf("missing parameter: role")
	}
	if role.Name == "" {
		return fmt.Errorf("missing parameter: name")
	}
	return nil
}

func (x *Create) Save(ctx context.Context, req *models.PostV1RolesJSONRequestBody) (*models.Role, error) {
	if err := x.validate(req); err != nil {
		return nil, err
	}
	res := &models.Role{
		Id:          x.identifier.UUIDv4(),
		Name:        strings.ToLower(req.Name),
		Description: req.Description,
	}
	if err := x.repository.InsertRole(ctx, res); err != nil {
		return nil, err
	}
	return res, nil
}
