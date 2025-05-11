package brokers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Publisher[T any] struct {
	client *redis.Client
	stream string
}

func NewPublisher[T any](client *redis.Client, stream string) *Publisher[T] {
	return &Publisher[T]{
		client: client,
		stream: stream,
	}
}

func (x *Publisher[T]) PublishMessage(ctx context.Context, payload T) (string, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to marshal payload: %w", err)
	}

	id, err := x.client.XAdd(ctx, &redis.XAddArgs{
		Stream: x.stream,
		Values: map[string]interface{}{
			"data": jsonData,
		},
	}).Result()
	if err != nil {
		return "", fmt.Errorf("failed to XAdd: %w", err)
	}
	return id, nil
}
