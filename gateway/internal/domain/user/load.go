package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/what-da-flac/wtf/gateway/internal/interfaces"
	"github.com/what-da-flac/wtf/openapi/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Load struct {
	repository interfaces.Repository
}

func NewLoad(repository interfaces.Repository) *Load {
	return &Load{
		repository: repository,
	}
}

func (x *Load) validate(req *models.User) error {
	r := req
	if r == nil {
		return fmt.Errorf("missing parameter: req")
	}
	if r.Email == "" && r.Id == "" {
		return fmt.Errorf("missing parameter: id or email")
	}
	return nil
}

func (x *Load) Load(ctx context.Context, req *models.User) (*models.User, error) {
	var id, email *string
	// validate incoming payload
	if err := x.validate(req); err != nil {
		return nil, err
	}
	if req.Id != "" {
		id = &req.Id
	}
	if req.Email != "" {
		email = &req.Email
	}
	res, err := x.repository.SelectUser(ctx, id, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, err
	}
	if res.IsDeleted {
		return nil, status.Error(codes.NotFound, "user not found")
	}
	return res, nil
}
