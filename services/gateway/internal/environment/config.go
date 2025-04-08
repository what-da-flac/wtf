package environment

import (
	"time"

	"github.com/spf13/viper"
	"github.com/what-da-flac/wtf/go-common/env"
)

const (
	envVarAPIUrlPrefix       = "API_URL_PREFIX"
	envVarCorsAllowedHeaders = "CORS_ALLOWED_HEADERS"
	envVarHeaderTimeout      = "HEADER_TIMEOUT"
)

type Config struct {
	env.Config

	APIUrlPrefix          string
	CorsAllowedHeaders    string
	ForceProfileCompleted bool
	HeaderTimeout         time.Duration
	SourceURL             string
}

func New() *Config {
	globalConfig := env.New()
	c := &Config{
		APIUrlPrefix:       viper.GetString(envVarAPIUrlPrefix),
		CorsAllowedHeaders: viper.GetString(envVarCorsAllowedHeaders),
		Config:             *globalConfig,
		HeaderTimeout:      viper.GetDuration(envVarHeaderTimeout),
	}
	return c
}
