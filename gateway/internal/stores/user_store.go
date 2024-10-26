package stores

import (
	"context"

	"github.com/what-da-flac/wtf/go-common/imodel"
)

type UserStore struct {
}

func NewUserStore() *UserStore {
	return &UserStore{}
}

func (x *UserStore) IsEnabled(ctx context.Context, user *imodel.User) (bool, error) {
	return user.Id != "", nil
}
