package environment

import "github.com/spf13/viper"

type Gateways struct {
	FileHost     string
	EmailHost    string
	ProfileHost  string
	TemplateHost string
	UserHost     string
}

const (
	envVarFileServiceHost     = "FILE_MANAGER_SERVICE_HOST"
	envVarEmailServiceHost    = "EMAIL_SERVICE_HOST"
	envVarTemplateServiceHost = "TEMPLATE_SERVICE_HOST"
	envVarProfileServiceHost  = "PROFILE_SERVICE_HOST"
	envVarUserServiceHost     = "USER_SERVICE_HOST"
)

func newGateways() Gateways {
	return Gateways{
		FileHost:     viper.GetString(envVarFileServiceHost),
		EmailHost:    viper.GetString(envVarEmailServiceHost),
		ProfileHost:  viper.GetString(envVarProfileServiceHost),
		TemplateHost: viper.GetString(envVarTemplateServiceHost),
		UserHost:     viper.GetString(envVarUserServiceHost),
	}
}
