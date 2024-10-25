package ifaces

import (
	"context"

	"github.com/what-da-flac/wtf/go-common/imodel"
)

type UserStore interface {
	IsEnabled(ctx context.Context, user *imodel.User) (bool, error)
}
