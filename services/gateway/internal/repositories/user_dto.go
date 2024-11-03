package repositories

import (
	"time"

	"github.com/what-da-flac/wtf/openapi/models"

	"github.com/jinzhu/copier"
)

type UserDto struct {
	Id        string
	Created   time.Time
	Email     string
	Name      string
	ImageUrl  *string
	LastLogin *time.Time
	IsDeleted bool
}

func (x *UserDto) toProto() *models.User {
	res := &models.User{}
	if err := copier.Copy(res, x); err != nil {
		return nil
	}
	res.Created = x.Created
	if val := x.LastLogin; val != nil {
		res.LastLogin = *val
	}
	res.Image = x.ImageUrl
	return res
}

func (x *UserDto) Fields() []string {
	return []string{"id", "email", "name", "created", "image_url", "last_login", "is_deleted"}
}

func (x *UserDto) TableName() string { return "users" }

func fromProtoUser(user *models.User) *UserDto {
	res := &UserDto{}
	if err := copier.Copy(res, user); err != nil {
		return nil
	}
	res.Created = user.Created
	res.LastLogin = &user.LastLogin
	res.ImageUrl = user.Image
	return res
}
