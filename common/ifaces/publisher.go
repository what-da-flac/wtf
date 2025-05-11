package ifaces

import "golang.org/x/net/context"

//go:generate moq -out ../mocks/publisher.go -pkg mocks . Publisher
type Publisher[T any] interface {
	PublishMessage(ctx context.Context, payload T) (string, error)
}
