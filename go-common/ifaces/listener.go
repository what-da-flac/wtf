package ifaces

type MessageHandlerFn func(msg []byte) (ack AckType, err error)

// Listener implements a message broker listener.
type Listener interface {
	Close() error
	ListenAsync(fn MessageHandlerFn)
}

type AckType int

const (
	// MessageAcknowledge marks message as processed and is deleted from queue.
	MessageAcknowledge AckType = iota

	// MessageReject marks message as unacknowledged and is deleted from queue.
	MessageReject AckType = iota

	// MessageRequeue does not mark message and sends to queue for next iteration.
	MessageRequeue AckType = iota
)
