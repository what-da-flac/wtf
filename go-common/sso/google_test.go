package sso

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"google.golang.org/api/idtoken"
)

func TestGoogleToken_parsePayload(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		args    args
		want    *idtoken.Payload
		wantErr bool
	}{
		{
			name: "happy path",
			args: args{
				token: `eyJhbGciOiJSUzI1NiIsImtpZCI6IjZmNzI1NDEwMWY1NmU0MWNmMzVjOTkyNmRlODRhMmQ1NTJiNGM2ZjEiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL2FjY291bnRzLmdvb2dsZS5jb20iLCJhenAiOiI5NTE3OTk5NzE2MTctZ2gxZXNlbmM4bWNzODI5OGdzMWNxb3NhbTlpNGhmZGMuYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJhdWQiOiI5NTE3OTk5NzE2MTctZ2gxZXNlbmM4bWNzODI5OGdzMWNxb3NhbTlpNGhmZGMuYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJzdWIiOiIxMTYzODY0NDg0MDg2MjE1NDc2MzkiLCJlbWFpbCI6Im1hdXJpY2lvLmxleXphb2xhQGdtYWlsLmNvbSIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJuYmYiOjE2OTU0OTIxMzcsIm5hbWUiOiJNYXVyaWNpbyBMZXl6YW9sYSIsInBpY3R1cmUiOiJodHRwczovL2xoMy5nb29nbGV1c2VyY29udGVudC5jb20vYS9BQ2c4b2NKZ25UZEVPeWtqZzFnb25raEZfdmR5V1lZcDdtSnBNU3I2SkszS3gxTjZudHM2PXM5Ni1jIiwiZ2l2ZW5fbmFtZSI6Ik1hdXJpY2lvIiwiZmFtaWx5X25hbWUiOiJMZXl6YW9sYSIsImxvY2FsZSI6ImVuIiwiaWF0IjoxNjk1NDkyNDM3LCJleHAiOjE2OTU0OTYwMzcsImp0aSI6IjZiOGYyNGQzODBmOWQ5ZTE2OTRjZjJiOWM4ZjE4NzMwMzYzZjg4ODYifQ.Kj9eVvoLzHSOPCoCEYotyjVQ5U3bC3JCuipDQ7GFKLPX65IjaXpvmi-HOwgpfkOOjrRh_QyW2dVW1-RqGtpEK9_Zcz3qelZhdS0FrMBLckve-8SdVOgBV-5kqwF8ZR6dhVGrj0npw7h9f-4fX1jDLOpD3X_B3BqzNYC8RXSY11AmrdYvukr_VHX5t9NUwVk0EvYoEHRRvXS6SCP66WxEQpt8YoPMIotqmvfuJ9clsa5lOsIMw7DCV_Br6E9oZz168qsejO-Dpih72gZmipL6DgkZ2ufxavlLaqGkhe4IInQkzoLfY2abB9jrrm_xkUxYc09KCxF5AY6kP-LdgtlQzA`,
			},
			wantErr: false,
			want: &idtoken.Payload{
				Issuer:   "https://accounts.google.com",
				Audience: "951799971617-gh1esenc8mcs8298gs1cqosam9i4hfdc.apps.googleusercontent.com",
				Expires:  1695496037,
				IssuedAt: 1695492437,
				Subject:  "116386448408621547639",
				Claims: map[string]interface{}{
					"aud":            "951799971617-gh1esenc8mcs8298gs1cqosam9i4hfdc.apps.googleusercontent.com",
					"azp":            "951799971617-gh1esenc8mcs8298gs1cqosam9i4hfdc.apps.googleusercontent.com",
					"email":          "mauricio.leyzaola@gmail.com",
					"email_verified": true,
					"exp":            1695496037,
					"family_name":    "Leyzaola",
					"given_name":     "Mauricio",
					"iat":            1695492437,
					"iss":            "https://accounts.google.com",
					"jti":            "6b8f24d380f9d9e1694cf2b9c8f18730363f8886",
					"locale":         "en",
					"name":           "Mauricio Leyzaola",
					"nbf":            1695492137,
					"picture":        "https://lh3.googleusercontent.com/a/ACg8ocJgnTdEOykjg1gonkhF_vdyWYYp7mJpMSr6JK3Kx1N6nts6=s96-c",
					"sub":            "116386448408621547639",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewGoogle()
			got, err := x.parsePayload(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("parsePayload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for k, v := range tt.want.Claims {
				val, ok := got.Claims[k]
				if !ok {
					t.Errorf("missing key: %s", k)
					continue
				}
				assert.EqualValues(t, v, val)
				delete(got.Claims, k)
			}
			assert.Empty(t, got.Claims)
			tt.want.Claims = got.Claims
			assert.EqualValues(t, tt.want, got)
		})
	}
}

func Test_googleClaimsToUserClaims(t *testing.T) {
	type args struct {
		i idtoken.Payload
	}
	tests := []struct {
		name string
		args args
		want *UserClaims
	}{
		{
			name: "happy path",
			args: args{
				i: idtoken.Payload{
					Issuer:   "https://accounts.google.com",
					Audience: "951799971617-gh1esenc8mcs8298gs1cqosam9i4hfdc.apps.googleusercontent.com",
					Expires:  1695496037,
					IssuedAt: 1695492437,
					Subject:  "116386448408621547639",
					Claims: map[string]interface{}{
						"aud":            "951799971617-gh1esenc8mcs8298gs1cqosam9i4hfdc.apps.googleusercontent.com",
						"azp":            "951799971617-gh1esenc8mcs8298gs1cqosam9i4hfdc.apps.googleusercontent.com",
						"email":          "mauricio.leyzaola@gmail.com",
						"email_verified": true,
						"exp":            1695496037,
						"family_name":    "Leyzaola",
						"given_name":     "Mauricio",
						"iat":            1695492437,
						"iss":            "https://accounts.google.com",
						"jti":            "6b8f24d380f9d9e1694cf2b9c8f18730363f8886",
						"locale":         "en",
						"name":           "Mauricio Leyzaola",
						"nbf":            1695492137,
						"picture":        "https://lh3.googleusercontent.com/a/ACg8ocJgnTdEOykjg1gonkhF_vdyWYYp7mJpMSr6JK3Kx1N6nts6=s96-c",
						"sub":            "116386448408621547639",
					},
				},
			},
			want: &UserClaims{
				Issuer:        "https://accounts.google.com",
				Audience:      "951799971617-gh1esenc8mcs8298gs1cqosam9i4hfdc.apps.googleusercontent.com",
				Expires:       1695496037,
				IssuedAt:      1695492437,
				Subject:       "116386448408621547639",
				Email:         "mauricio.leyzaola@gmail.com",
				EmailVerified: true,
				FamilyName:    "Leyzaola",
				GivenName:     "Mauricio",
				PictureURL:    "https://lh3.googleusercontent.com/a/ACg8ocJgnTdEOykjg1gonkhF_vdyWYYp7mJpMSr6JK3Kx1N6nts6=s96-c",
				Locale:        "en",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, googleClaimsToUserClaims(tt.args.i), "googleClaimsToUserClaims(%v)", tt.args.i)
		})
	}
}
