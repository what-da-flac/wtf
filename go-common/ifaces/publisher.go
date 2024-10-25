package ifaces

import (
	"context"
)

// Publisher implements a message broker publisher.
type Publisher interface {
	Build() error
	Close() error
	Name() string
	Publish(ctx context.Context, data []byte) error
	WithConnection(protocol, username, password, hostname, port string) Publisher
	WithName(name string) Publisher
}
