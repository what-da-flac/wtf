package env

import "github.com/spf13/viper"

const envVarRabbitMQUrl = "RABBITMQ_URL"

type RabbitMQ struct {
	URL string
}

func newRabbitMQ() RabbitMQ {
	return RabbitMQ{
		URL: viper.GetString(envVarRabbitMQUrl),
	}
}
