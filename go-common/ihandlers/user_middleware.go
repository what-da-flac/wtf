package ihandlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/what-da-flac/wtf/go-common/ifaces"
)

// UserMiddleware checks for valid JWT in the headers and parses a user struct.
// If all goes well, user struct gets injects into context for
// next middleware to be used.
func UserMiddleware(
	urlPrefix string,
	decider ifaces.EndpointDecider,
	userStore ifaces.UserStore,
) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			path, method := r.URL.Path, r.Method
			// remove the urlPrefix from the path
			path = strings.TrimPrefix(path, urlPrefix)
			// check if this path/method does not require any permission
			if !decider.IsSecured(path, method) {
				h.ServeHTTP(w, r)
				return
			}
			ctx := r.Context()
			user := UserFromContext(ctx)
			if user == nil {
				WriteResponse(w, http.StatusNotFound, nil, fmt.Errorf("user not found in conext"))
				return
			}
			// check user is enabled
			if ok, err := userStore.IsEnabled(ctx, user); err != nil {
				WriteResponse(w, http.StatusInternalServerError, nil, err)
				return
			} else if !ok {
				WriteResponse(w, http.StatusUnauthorized, nil, err)
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}
