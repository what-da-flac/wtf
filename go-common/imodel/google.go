package imodel

// GoogleUserInfo is the data structure that Google returns when validating a JWT token.
type GoogleUserInfo struct {
	Id            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

func (x *GoogleUserInfo) ToUser() *User {
	return &User{
		Email: x.Email,
		Id:    x.Id,
		Image: &x.Picture,
		Name:  x.Name,
	}
}
