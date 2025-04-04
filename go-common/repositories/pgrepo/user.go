package pgrepo

import (
	"context"
	"fmt"

	"github.com/what-da-flac/wtf/openapi/gen/golang"
)

func (x *PgRepo) InsertUser(_ context.Context, user *golang.User) error {
	db := x.GORM()
	row := fromProtoUser(user)
	return db.Create(row).Error
}

func (x *PgRepo) UpdateUser(cxt context.Context, user *golang.User) error {
	db := x.GORM()
	row := fromProtoUser(user)
	return db.Updates(row).Error
}

func (x *PgRepo) SelectUser(cxt context.Context, id, email *string) (*golang.User, error) {
	db := x.GORM()
	row := &UserDto{}
	where := &UserDto{}
	switch {
	case id != nil:
		where.Id = *id
	case email != nil:
		where.Email = *email
	default:
		return nil, fmt.Errorf("missing id or email")
	}
	db = db.Where(where).First(row)
	if err := db.Error; err != nil {
		return nil, err
	}
	return row.toProto(), nil
}

func (x *PgRepo) ListUsers(_ context.Context, req *golang.UserListParams) ([]*golang.User, error) {
	var (
		result []*golang.User
		rows   []*UserDto
	)
	instance := &UserDto{}
	db := x.GORM()
	db = db.Model(instance).Select(instance.Fields())
	params := req
	if len(params.Emails) != 0 {
		db = db.Where("(email) in ?", params.Emails)
	}
	if len(params.Ids) != 0 {
		db = db.Where("(id) in ?", params.Ids)
	}
	if val := params.EmailMatch; val != "" {
		val = "%" + val + "%"
		db = db.Where("email ILIKE ? OR name ILIKE ?", val, val)
	}
	// we need to use a map because boolean false won't be evaluated by gorm
	db = db.Where("is_deleted = ?", params.OnlyDeleted)
	if val := params.Limit; val != nil {
		db = db.Limit(*val)
	} else {
		db = db.Limit(x.defaultLimit)
	}
	if params.Offset != 0 {
		db = db.Offset(params.Offset)
	}
	db = db.Find(&rows)
	if err := db.Error; err != nil {
		return nil, err
	}
	for _, row := range rows {
		result = append(result, row.toProto())
	}
	return result, nil
}
