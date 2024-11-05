package environment

import (
	"time"

	"github.com/spf13/viper"
	"github.com/what-da-flac/wtf/go-common/environment"
)

const (
	envVarAPIUrlPrefix       = "API_URL_PREFIX"
	envVarCorsAllowedHeaders = "CORS_ALLOWED_HEADERS"
	envVarHeaderTimeout      = "HEADER_TIMEOUT"
)

type Config struct {
	environment.Config

	APIUrlPrefix          string
	CorsAllowedHeaders    string
	ForceProfileCompleted bool
	HeaderTimeout         time.Duration
	SourceURL             string
}

func New() *Config {
	globalConfig := environment.New()
	return &Config{
		APIUrlPrefix:       viper.GetString(envVarAPIUrlPrefix),
		CorsAllowedHeaders: viper.GetString(envVarCorsAllowedHeaders),
		Config:             *globalConfig,
		HeaderTimeout:      viper.GetDuration(envVarHeaderTimeout),
	}
}
