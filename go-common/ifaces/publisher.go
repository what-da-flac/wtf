package ifaces

// Publisher implements a message broker publisher.
type Publisher interface {
	Build() error
	Close() error
	Publish(data []byte) error
}
