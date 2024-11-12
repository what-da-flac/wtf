package rabbits

import (
	"fmt"
	"time"

	"github.com/what-da-flac/wtf/go-common/env"

	"github.com/rabbitmq/amqp091-go"
	"github.com/what-da-flac/wtf/go-common/ifaces"
)

// Listener implements Listener interface using RabbitMQ as message broker.
type Listener struct {
	interval time.Duration
	logger   ifaces.Logger
	name     string
	uri      string

	conn *amqp091.Connection
	ch   *amqp091.Channel
}

// NewListener returns an instance of RabbitMQ message listener.
// name: unique name of queue.
// uri: fully qualified rabbitMQ url to connect at.
// interval: the amount of time listener waits for next message.
func NewListener(logger ifaces.Logger, name env.Names, uri string, interval time.Duration) *Listener {
	return &Listener{
		logger:   logger,
		name:     name.String(),
		uri:      uri,
		interval: interval,
	}
}

// Close closes all resources.
func (x *Listener) Close() error {
	if err := x.ch.Close(); err != nil {
		return err
	}
	return x.conn.Close()
}

func (x *Listener) connect() error {
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

// ListenAsync is the exposed listener.
func (x *Listener) ListenAsync(fn ifaces.MessageHandlerFn) {
	go func() {
		if err := x.listen(fn); err != nil {
			panic(err)
		}
	}()
}

// listen should be called from a goroutine because it blocks execution.
// fn: callback function to execute on each message received.
// Returns error if connection was not successful.
func (x *Listener) listen(fn ifaces.MessageHandlerFn) error {
	// Connect to RabbitMQ
	if err := x.connect(); err != nil {
		return err
	}
	// Declare a queue
	queue, err := x.createQueue()
	if err != nil {
		x.logger.Errorf("Failed to create queue: %v", err)
		return err
	}

	// Start consuming messages
	msgs, err := x.ch.Consume(
		queue.Name, // queue
		"",         // consumer
		false,      // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		return fmt.Errorf("failed to register a consumer: %w", err)
	}

	x.logger.Infof("listening for messages on queue %s every %v", x.name, x.interval)

	// Process messages with the given interval
	for msg := range msgs {
		x.processMessage(fn, msg)
		// Wait for the specified interval before processing the next message
		time.Sleep(x.interval)
	}

	return nil
}

func (x *Listener) createQueue() (*amqp091.Queue, error) {
	q, err := x.ch.QueueDeclare(
		x.name, // name
		true,   // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare a queue: %w", err)
	}
	return &q, nil
}

func (x *Listener) processMessage(fn ifaces.MessageHandlerFn, msg amqp091.Delivery) {
	// Process the message here
	ackType, err := fn(msg.Body)
	if err != nil {
		x.logger.Errorf("failed to process message: %v", err)
	}
	switch ackType {
	case ifaces.MessageAcknowledge:
		if err = msg.Ack(false); err != nil {
			x.logger.Errorf("failed to ack message: %v", err)
		}
	case ifaces.MessageReject:
		if err = msg.Nack(false, false); err != nil {
			x.logger.Errorf("failed to reject message: %v", err)
		}
	case ifaces.MessageRequeue:
		if err = msg.Nack(false, true); err != nil {
			x.logger.Errorf("failed to requeue message: %v", err)
		}
	}
}
