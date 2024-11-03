package interfaces

import "context"

//go:generate moq -out ../../mocks/message_listener.go -pkg mocks . MessageListener
type MessageListener interface {
	Name() string
	Poll(ctx context.Context) error
}

type MessageReceiverFn func(body string) error
