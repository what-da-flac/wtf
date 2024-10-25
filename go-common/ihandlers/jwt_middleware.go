package ihandlers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/what-da-flac/wtf/go-common/ifaces"
	"github.com/what-da-flac/wtf/go-common/imodel"
)

// JWTMiddleware checks for valid JWT in the headers and parses a user struct.
// If all goes well, user struct gets injects into context for
// next middleware to be used.
func JWTMiddleware(
	urlPrefix string,
	timeout time.Duration,
	decider ifaces.EndpointDecider,
	userStore ifaces.UserStore,
	validator ifaces.TokenValidator,
	instantiateUserFn func() any, // expects a new user instance
	convertUserFn func(any) (*imodel.User, error), // post process the user after token has been validated
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

			// read jwt token from request header and set the new context
			authToken := r.Header.Get(authKey)
			ssoUser := instantiateUserFn()
			if err := validator.Validate(authToken, timeout, ssoUser); err != nil {
				err = fmt.Errorf("invalid google user: %w", err)
				WriteResponse(w, http.StatusUnauthorized, nil, err)
				return
			}
			newUser, err := convertUserFn(ssoUser)
			if err != nil {
				WriteResponse(w, http.StatusInternalServerError, nil, err)
				return
			}
			ctx := SetUserInContext(r.Context(), newUser)

			// check user is enabled
			if ok, err := userStore.IsEnabled(ctx, newUser); err != nil {
				WriteResponse(w, http.StatusForbidden, nil, err)
				return
			} else if !ok {
				WriteResponse(w, http.StatusUnauthorized, nil, err)
				return
			}
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
