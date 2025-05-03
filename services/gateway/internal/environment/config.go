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
	envVarPort               = "GATEWAY_PORT"
)

type Config struct {
	env.Config

	APIUrlPrefix          string
	CorsAllowedHeaders    string
	ForceProfileCompleted bool
	HeaderTimeout         time.Duration
	Port                  string
	SourceURL             string
}

func New() *Config {
	genv := env.New()
	c := &Config{
		APIUrlPrefix:       viper.GetString(envVarAPIUrlPrefix),
		CorsAllowedHeaders: viper.GetString(envVarCorsAllowedHeaders),
		Config:             *genv,
		HeaderTimeout:      viper.GetDuration(envVarHeaderTimeout),
		Port:               viper.GetString(envVarPort),
	}
	return c
}
