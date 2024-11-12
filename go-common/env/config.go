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
	AWS         AWS
	DB          DB
	Env         string
	Google      Google
	LogLevel    string
	Port        string
	RabbitMQ    RabbitMQ
	Sentry      Sentry
	ServiceName string
	Volumes     Volumes
}

func New() *Config {
	const defaultLogLevel = "INFO"
	viper.AutomaticEnv()
	c := &Config{
		AWS:         newAWS(),
		DB:          newDB(),
		Env:         viper.GetString(envVarEnv),
		Google:      newGoogle(),
		LogLevel:    viper.GetString(envLogLevel),
		Port:        viper.GetString(envVarPort),
		RabbitMQ:    newRabbitMQ(),
		Sentry:      newSentry(),
		ServiceName: viper.GetString(envVarServiceName),
		Volumes:     newVolumes(),
	}
	if c.LogLevel == "" {
		c.LogLevel = defaultLogLevel
	}
	return c
}
