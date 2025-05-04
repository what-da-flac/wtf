package env

import (
	"github.com/spf13/viper"
)

const (
	envVarEnv         = "ENV"
	envVarServiceName = "SERVICE_NAME"
	envVarPort        = "PORT"
	envLogLevel       = "LOG_LEVEL"
)

type Config struct {
	DB          DB
	Env         string
	LogLevel    string
	Paths       *Path
	ServiceName string
}

func New() *Config {
	const defaultLogLevel = "INFO"
	viper.AutomaticEnv()
	c := &Config{
		DB:          newDB(),
		Env:         viper.GetString(envVarEnv),
		LogLevel:    viper.GetString(envLogLevel),
		Paths:       NewPathsFromEnvironment(),
		ServiceName: viper.GetString(envVarServiceName),
	}
	if c.LogLevel == "" {
		c.LogLevel = defaultLogLevel
	}
	return c
}
