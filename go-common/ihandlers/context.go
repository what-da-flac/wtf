package ihandlers

import (
	"context"

	"github.com/what-da-flac/wtf/go-common/imodel"
)

func setContextValue(ctx context.Context, key keyStr, v any) context.Context {
	//nolint:staticcheck
	return context.WithValue(ctx, key, v)
}

// UserFromContext reads the user provided in the request without doing a db trip.
// Will return nil if there is no user in the context of the request.
func UserFromContext(ctx context.Context) *imodel.User {
	if val, ok := ctx.Value(userKey).(*imodel.User); ok {
		return val
	}
	return nil
}

func SetUserInContext(ctx context.Context, user any) context.Context {
	return setContextValue(ctx, userKey, user)
}

func RolesFromContext(ctx context.Context) []string {
	val, ok := ctx.Value(rolesKey).([]string)
	if !ok {
		return nil
	}
	return val
}
