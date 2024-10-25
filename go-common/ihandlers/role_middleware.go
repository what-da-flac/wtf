package ihandlers

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/what-da-flac/wtf/go-common/ifaces"
	"github.com/what-da-flac/wtf/go-common/imodel"
)

// RoleMiddleware checks for user's role membership and returns a 401 if user is not member of any role.
// If all goes well, roles array will be injected into context for next middlewares to use it.
func RoleMiddleware(
	urlPrefix string,
	decider ifaces.EndpointDecider,
	readUserRolesFn func(ctx context.Context, user *imodel.User) ([]string, error),
) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			path, method := r.URL.Path, r.Method
			// remove the urlPrefix from the path
			path = strings.TrimPrefix(path, urlPrefix)
			// check endpoint is secured
			if !decider.IsSecured(path, method) {
				h.ServeHTTP(w, r)
				return
			}
			ctx := r.Context()

			// grab user from context
			user := UserFromContext(ctx)
			if user == nil {
				WriteResponse(w, http.StatusNotImplemented, nil, fmt.Errorf("User not found in context"))
				return
			}

			// check user has role membership
			roles, err := readUserRolesFn(ctx, user)
			if err != nil {
				WriteResponse(w, http.StatusNotImplemented, nil, err)
				return
			}

			// set the new context with roles
			ctx = setContextValue(ctx, rolesKey, roles)

			// check if decider allows to continue
			if !decider.Allow(path, method, roles...) {
				WriteResponse(w, http.StatusForbidden, nil, fmt.Errorf("forbidden"))
				return
			}
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
