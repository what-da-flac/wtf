package sso

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"google.golang.org/api/idtoken"
)

// Google is a placeholder for sso operations implemented with Google Auth0.
type Google struct{}

func NewGoogle() *Google {
	return &Google{}
}

func (x *Google) Validate(ctx context.Context, token string) (*UserClaims, error) {
	payload, err := idtoken.Validate(ctx, token, "")
	if err != nil {
		return nil, err
	}
	return googleClaimsToUserClaims(*payload), nil
}

//nolint:unused
func (x *Google) parsePayload(token string) (*idtoken.Payload, error) {
	return idtoken.ParsePayload(token)
}

func googleClaimsToUserClaims(i idtoken.Payload) *UserClaims {
	return &UserClaims{
		Issuer:        i.Issuer,
		Audience:      i.Audience,
		Expires:       i.Expires,
		IssuedAt:      i.IssuedAt,
		Subject:       i.Subject,
		Email:         InterfaceToString(i.Claims["email"]),
		EmailVerified: InterfaceToBool(i.Claims["email_verified"]),
		FamilyName:    InterfaceToString(i.Claims["family_name"]),
		GivenName:     InterfaceToString(i.Claims["given_name"]),
		PictureURL:    InterfaceToString(i.Claims["picture"]),
		Locale:        InterfaceToString(i.Claims["locale"]),
	}
}

func (x *Google) Type() ProviderType { return GoogleProvider }

// ValidateJWTToken makes a request to google apis to check the user is a valid one.
// Parameter v must be a pointer receiver, compatible with GoogleUserInfo.
func ValidateJWTToken(accessToken string, timeout time.Duration, v interface{}) error {
	const (
		accessTokenParamName = "access_token"
		googleEndpoint       = "https://www.googleapis.com/oauth2/v1/userinfo"
	)
	client := &http.Client{
		Timeout: timeout,
	}
	params := url.Values{}
	params.Set(accessTokenParamName, accessToken)
	uri := googleEndpoint + "?" + params.Encode()
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return err
	}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	if status := res.StatusCode; status == http.StatusUnauthorized {
		return fmt.Errorf("unauthorized")
	}
	defer func() { _ = res.Body.Close() }()
	return json.NewDecoder(res.Body).Decode(v)
}

func InterfaceToString(i interface{}) string {
	if val, ok := i.(string); ok {
		return val
	}
	return ""
}

func InterfaceToBool(i interface{}) bool {
	if val, ok := i.(bool); ok {
		return val
	}
	return false
}
