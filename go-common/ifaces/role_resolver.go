package ifaces

import (
	"context"

	"github.com/what-da-flac/wtf/go-common/imodel"
)

// RoleStore reads from any data source the assigned roles for a given user, based on its id.
type RoleStore interface {
	Roles(ctx context.Context, user *imodel.User) ([]string, error)
}
