package environment

import "github.com/spf13/viper"

const (
	envVarRabbitMQHostname = "RABBITMQ_HOSTNAME"
	envVarRabbitMQPassword = "RABBITMQ_PASSWORD"
	envVarRabbitMQPort     = "RABBITMQ_PORT"
	envVarRabbitMQProtocol = "RABBITMQ_PROTOCOL"
	envVarRabbitMQUsername = "RABBITMQ_USERNAME"
	envVarRabbitMQUrl      = "RABBITMQ_URL"
)

type RabbitMQ struct {
	Hostname string
	Password string
	Port     string
	Protocol string
	Username string
	URL      string
}

func newRabbitMQ() RabbitMQ {
	return RabbitMQ{
		Hostname: viper.GetString(envVarRabbitMQHostname),
		Password: viper.GetString(envVarRabbitMQPassword),
		Port:     viper.GetString(envVarRabbitMQPort),
		Protocol: viper.GetString(envVarRabbitMQProtocol),
		Username: viper.GetString(envVarRabbitMQUsername),
		URL:      viper.GetString(envVarRabbitMQUrl),
	}
}
