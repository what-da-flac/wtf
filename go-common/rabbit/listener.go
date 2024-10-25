package rabbit

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/what-da-flac/wtf/go-common/ifaces"
)

const (
	defaultExchangeKind = "direct"
)

type Listener struct {
	// rabbitmq objects we need to attach
	conn     *amqp.Connection
	channel  *amqp.Channel
	messages <-chan amqp.Delivery

	ackErrHandler    ifaces.ErrHandlerFn
	handler          ifaces.MessageHandlerFn
	connectionParams *connectionParams

	// delegate naming convention
	namer *namer
	name  string
}

func NewListener() *Listener {
	return &Listener{}
}

func (x *Listener) WithConnection(protocol, username, password, hostname, port string) ifaces.Listener {
	x.connectionParams = &connectionParams{
		Protocol: protocol,
		Username: username,
		Password: password,
		Hostname: hostname,
		Port:     port,
	}
	return x
}

func (x *Listener) WithHandler(handler ifaces.MessageHandlerFn) ifaces.Listener {
	x.handler = handler
	return x
}

func (x *Listener) WithName(name string) ifaces.Listener {
	x.name = name
	x.namer = newNamer(name)
	return x
}

func (x *Listener) Name() string { return x.name }

func (x *Listener) WithAckErrorHandler(ackErrHandler ifaces.ErrHandlerFn) ifaces.Listener {
	x.ackErrHandler = ackErrHandler
	return x
}

func (x *Listener) validate() error {
	if x.name == "" {
		return fmt.Errorf("missing parameter: name")
	}
	if x.handler == nil {
		return fmt.Errorf("missing parameter: handler")
	}
	if x.ackErrHandler == nil {
		return fmt.Errorf("missing parameter: ackErrHandler")
	}
	if x.connectionParams == nil {
		return fmt.Errorf("missing parameter: connectionParams")
	}
	return nil
}

func (x *Listener) Build() error {
	// sanity check
	if err := x.validate(); err != nil {
		return err
	}
	// rabbitmq connection
	p := x.connectionParams
	mqConn, err := Connect(p.Protocol, p.Username, p.Password, p.Hostname, p.Port)
	if err != nil {
		return err
	}

	// main channel
	channel, err := mqConn.Channel()
	if err != nil {
		return err
	}
	// bind channel with exchanges and queues
	queue, err := x.bind(channel)
	if err != nil {
		return err
	}
	// assign messages, so we can dispatch them
	messages, err := x.assignMessages(channel, *queue)
	if err != nil {
		return err
	}
	x.conn = mqConn
	x.channel = channel
	x.messages = messages

	return nil
}

// bind is internal method to create exchanges and queues, also to link them together,
// or implement more complex scenarios in the future.
func (x *Listener) bind(channel *amqp.Channel) (*amqp.Queue, error) {
	// main exchange
	if err := x.declareMainExchange(channel); err != nil {
		return nil, err
	}
	// main queue
	queue, err := x.declareMainQueue(channel)
	if err != nil {
		return nil, err
	}
	// bind queue to exchange
	if err := x.bindMainQueueToExchange(channel); err != nil {
		return nil, err
	}
	return &queue, nil
}

func (x *Listener) assignMessages(channel *amqp.Channel, queue amqp.Queue) (<-chan amqp.Delivery, error) {
	const defaultConsumerName = "consumer"
	return channel.Consume(
		queue.Name,
		defaultConsumerName,
		false,
		false,
		false,
		false,
		nil,
	)
}

func (x *Listener) bindMainQueueToExchange(channel *amqp.Channel) error {
	return channel.QueueBind(
		x.namer.queueName(),
		x.namer.keyName(),
		x.namer.exchangeName(),
		false,
		nil,
	)
}

func (x *Listener) declareMainQueue(channel *amqp.Channel) (amqp.Queue, error) {
	return channel.QueueDeclare(
		x.namer.queueName(),
		true,
		false,
		false,
		false,
		nil,
	)
}

func (x *Listener) declareMainExchange(channel *amqp.Channel) error {
	return channel.ExchangeDeclare(
		x.namer.exchangeName(),
		defaultExchangeKind,
		true,
		false,
		false,
		false,
		nil,
	)
}

func (x *Listener) Close() error {
	if err := x.channel.Close(); err != nil {
		return err
	}
	return x.conn.Close()
}

func (x *Listener) ListenAsync() {
	go func() {
		for m := range x.messages {
			// receive the ack type from handler
			ackType, err := x.handler(m.Body)
			if err != nil {
				// let caller deal with error
				x.ackErrHandler(m.Body, err)
			}
			err = nil
			switch ackType {
			case ifaces.MessageAcknowledge:
				err = m.Ack(false)
			case ifaces.MessageReject:
				err = m.Reject(false)
			case ifaces.MessageRequeue:
				err = m.Nack(false, false)
			}
			if err != nil {
				// let caller deal with error
				x.ackErrHandler(m.Body, err)
			}
		}
	}()
}
