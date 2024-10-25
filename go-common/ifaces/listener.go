package ifaces

type MessageHandlerFn func(msg []byte) (ack AckType, err error)

type ErrHandlerFn func(msg []byte, err error)

// Listener implements a message broker listener.
type Listener interface {
	Build() error
	Close() error
	ListenAsync()
	Name() string
	WithAckErrorHandler(errHandler ErrHandlerFn) Listener
	WithConnection(protocol, username, password, hostname, port string) Listener
	WithHandler(handler MessageHandlerFn) Listener
	WithName(name string) Listener
}

type AckType int

const (
	MessageAcknowledge AckType = iota
	MessageReject      AckType = iota
	MessageRequeue     AckType = iota
)
