package imodel

import (
	"context"
	"time"
)

type EndpointDecider interface {
	// Allow returns true if based on path, method and roles provided, the user has access granted to move on.
	Allow(path, method string, roles ...string) bool

	// IsSecured returns true if the endpoint has been marked as insecure, in which case there is no need to process
	// additional authentication checks.
	IsSecured(path, method string) bool
}

type RoleStore interface {
	// Roles returns the role the user provided is member of.
	Roles(ctx context.Context, userId string) ([]string, error)
}

type TokenValidator interface {

	// Validate checks against a SSO provider, the token is valid and has not expired yet.
	Validate(token string, timeout time.Duration, v interface{}) error
}
