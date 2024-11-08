package rabbits

import (
	"fmt"

	"github.com/what-da-flac/wtf/go-common/env"

	"github.com/rabbitmq/amqp091-go"
	"github.com/what-da-flac/wtf/go-common/ifaces"
)

type Publisher struct {
	built  bool
	logger ifaces.Logger
	name   string
	uri    string

	conn  *amqp091.Connection
	ch    *amqp091.Channel
	queue *amqp091.Queue
}

func NewPublisher(logger ifaces.Logger, name env.QueueName, uri string) *Publisher {
	return &Publisher{
		built:  false,
		logger: logger,
		name:   name.String(),
		uri:    uri,
	}
}

func (x *Publisher) Publish(data []byte) error {
	if !x.built {
		return fmt.Errorf("publisher is not built")
	}
	return x.ch.Publish(
		"",
		x.name,
		false,
		false,
		amqp091.Publishing{
			Body: data,
		},
	)
}

// Close closes all resources.
func (x *Publisher) Close() error {
	if err := x.ch.Close(); err != nil {
		return err
	}
	return x.conn.Close()
}

func (x *Publisher) Build() error {
	if x.built {
		return nil
	}
	if x.name == "" {
		return fmt.Errorf("publisher name is empty")
	}
	if x.uri == "" {
		return fmt.Errorf("publisher uri is empty")
	}
	if x.logger == nil {
		return fmt.Errorf("publisher logger is nil")
	}
	if err := x.connect(); err != nil {
		return err
	}
	queue, err := x.ch.QueueDeclare(
		x.name,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	x.queue = &queue
	x.built = true
	return nil
}

func (x *Publisher) connect() error {
	conn, err := amqp091.Dial(x.uri)
	if err != nil {
		return err
	}
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	x.conn = conn
	x.ch = ch
	return nil
}
