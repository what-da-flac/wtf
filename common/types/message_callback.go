package types

type MessageCallback[T any] func(msg T) (ack bool, err error)
