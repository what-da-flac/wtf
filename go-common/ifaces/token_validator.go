package ifaces

import "time"

type TokenValidator interface {

	// Validate checks against a SSO provider, the token is valid and has not expired yet.
	Validate(token string, timeout time.Duration, v interface{}) error
}
