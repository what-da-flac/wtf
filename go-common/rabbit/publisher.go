package rabbit

import (
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/what-da-flac/wtf/go-common/ifaces"
)

type Publisher struct {
	conn    *amqp.Connection
	channel *amqp.Channel

	namer            *namer
	name             string
	connectionParams *connectionParams
}

func NewPublisher() *Publisher {
	return &Publisher{}
}

func (x *Publisher) validate() error {
	if x.name == "" {
		return fmt.Errorf("missing parameter: name")
	}
	if x.connectionParams == nil {
		return fmt.Errorf("missing parameter: connectionParams")
	}
	return nil
}

func (x *Publisher) WithConnection(protocol, username, password, hostname, port string) ifaces.Publisher {
	x.connectionParams = &connectionParams{
		Protocol: protocol,
		Username: username,
		Password: password,
		Hostname: hostname,
		Port:     port,
	}
	return x
}

func (x *Publisher) WithName(name string) ifaces.Publisher {
	x.name = name
	x.namer = newNamer(name)
	return x
}

func (x *Publisher) Publish(ctx context.Context, data []byte) error {
	msg := amqp.Publishing{
		Body: data,
	}
	return x.channel.PublishWithContext(
		ctx,
		x.namer.exchangeName(),
		x.namer.keyName(),
		false,
		false,
		msg,
	)
}

func (x *Publisher) Name() string { return x.name }

func (x *Publisher) Build() error {
	if err := x.validate(); err != nil {
		return err
	}
	p := x.connectionParams
	mqConn, err := Connect(p.Protocol, p.Username, p.Password, p.Hostname, p.Port)
	if err != nil {
		return err
	}
	channel, err := mqConn.Channel()
	if err != nil {
		return err
	}
	x.conn = mqConn
	x.channel = channel
	return nil
}

func (x *Publisher) Close() error {
	if err := x.channel.Close(); err != nil {
		return err
	}
	return x.conn.Close()
}
