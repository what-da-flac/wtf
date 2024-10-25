package sso

type ProviderType int

const (
	GoogleProvider ProviderType = iota
)

// UserClaims represents the claims from a JWT token.
// Only Google is implemented for the time being.
type UserClaims struct {
	Issuer   string
	Audience string
	Expires  int64
	IssuedAt int64
	Subject  string

	// custom fields not necessary present in claims
	Email         string
	EmailVerified bool
	FamilyName    string
	GivenName     string
	PictureURL    string
	Locale        string
}
