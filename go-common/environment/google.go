package environment

import "github.com/spf13/viper"

const (
	//nolint:gosec
	envVarGoogleApiKey       = "GOOGLE_API_KEY"
	envVarGoogleClientId     = "GOOGLE_CLIENT_ID"
	envVarGoogleClientSecret = "GOOGLE_CLIENT_SECRET"
)

type Google struct {
	ApiKey       string
	ClientId     string
	ClientSecret string
}

func newGoogle() Google {
	return Google{
		ApiKey:       viper.GetString(envVarGoogleApiKey),
		ClientId:     viper.GetString(envVarGoogleClientId),
		ClientSecret: viper.GetString(envVarGoogleClientSecret),
	}
}
