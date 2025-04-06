package environment

import (
	"errors"
	"strconv"
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
	Port                  string
	port                  int
	SourceURL             string
}

func New() (*Config, error) {
	globalConfig := env.New()
	c := &Config{
		APIUrlPrefix:       viper.GetString(envVarAPIUrlPrefix),
		CorsAllowedHeaders: viper.GetString(envVarCorsAllowedHeaders),
		Config:             *globalConfig,
		HeaderTimeout:      viper.GetDuration(envVarHeaderTimeout),
		port:               viper.GetInt("PORT"),
	}
	if err := c.validate(); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Config) validate() error {
	if c.port == 0 {
		return errors.New("port must be set")
	}
	c.Port = strconv.Itoa(c.port)
	return nil
}
