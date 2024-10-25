package sso

import (
	"fmt"
	"strings"
	"time"
)

type GoogleValidator struct{}

func NewGoogleValidator() *GoogleValidator {
	return &GoogleValidator{}
}

func (x *GoogleValidator) Validate(token string, timeout time.Duration, v interface{}) error {
	values := strings.Split(token, " ")
	if len(values) != 2 {
		return fmt.Errorf("cannot parse bearer token")
	}
	accessToken := values[1]
	return ValidateJWTToken(accessToken, timeout, v)
}
