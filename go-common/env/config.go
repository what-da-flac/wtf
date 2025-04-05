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
	Env         string
	LogLevel    string
	Port        string
	ServiceName string
}

func New() *Config {
	const defaultLogLevel = "INFO"
	viper.AutomaticEnv()
	c := &Config{
		Env:         viper.GetString(envVarEnv),
		LogLevel:    viper.GetString(envLogLevel),
		Port:        viper.GetString(envVarPort),
		ServiceName: viper.GetString(envVarServiceName),
	}
	if c.LogLevel == "" {
		c.LogLevel = defaultLogLevel
	}
	return c
}
