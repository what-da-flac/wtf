package brokers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/what-da-flac/wtf/go-common/types"
)

const (
	consumer         = "json_consumer"
	existingGroupErr = "BUSYGROUP Consumer Group name already exists"
)

// MessageWrapper contains the raw Redis Stream message and decoded payload of type T
type MessageWrapper[T any] struct {
	ID      string
	Payload T
	Raw     map[string]interface{}
}

type Subscriber[T any] struct {
	client *redis.Client
	group  string
	stream string
}

func NewSubscriber[T any](client *redis.Client, stream, group string) (*Subscriber[T], error) {
	var ctx = context.Background()
	// create group and ignore error if already exists
	if err := client.XGroupCreateMkStream(ctx, stream, group, "0").Err(); err != nil && err.Error() != existingGroupErr {
		return nil, fmt.Errorf("failed to create consumer group: %v", err)
	}
	x := &Subscriber[T]{
		client: client,
		group:  group,
		stream: stream,
	}

	return x, nil
}

func (x *Subscriber[T]) Listen(ctx context.Context, processMessageFn types.MessageCallback[T], errFn func(error)) {
	// clear any pending messages
	if err := x.clearPendingMessages(ctx, consumer); err != nil {
		errFn(err)
		return
	}
	// continuously read new messages
	for {
		messages, err := x.fetchNewMessages(ctx, consumer, 1, 2*time.Second)
		if err != nil {
			return
		}
		for _, msg := range messages {
			ok, err := processMessageFn(msg.Payload)
			if err != nil {
				errFn(err)
				continue
			}
			if !ok {
				continue
			}
			// acknowledge the message after processing
			if err := x.client.XAck(ctx, x.stream, x.group, msg.ID).Err(); err != nil {
				errFn(err)
			}
			// delete the message from the stream
			if err := x.client.XDel(ctx, x.stream, msg.ID).Err(); err != nil {
				errFn(err)
			}
		}
	}
}

func (x *Subscriber[T]) clearPendingMessages(ctx context.Context, consumer string) error {
	pendingMessages, err := x.fetchPendingMessages(ctx, consumer, 10)
	if err != nil {
		log.Fatalf("Error fetching pending messages: %v", err)
	}

	for _, msg := range pendingMessages {
		// acknowledge the message after processing
		if err := x.client.XAck(ctx, x.stream, x.group, msg.ID).Err(); err != nil {
			return err
		}
		// delete the message from the stream
		if err := x.client.XDel(ctx, x.stream, msg.ID).Err(); err != nil {
			return err
		}
	}
	return nil
}

func (x *Subscriber[T]) fetchPendingMessages(ctx context.Context, consumer string, count int64) ([]MessageWrapper[T], error) {
	// fetch pending entries for this consumer
	pending, err := x.client.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    x.group,
		Consumer: consumer,
		Streams:  []string{x.stream, "0"}, // "0" = read pending
		Count:    count,
		Block:    0, // wait if no messages
	}).Result()

	if err != nil {
		return nil, fmt.Errorf("failed to fetch pending: %w", err)
	}
	return x.parseStreamMessages(pending)
}

func (x *Subscriber[T]) fetchNewMessages(ctx context.Context, consumer string, count int64, blockDuration time.Duration) ([]MessageWrapper[T], error) {
	// Fetch new messages
	messages, err := x.client.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    x.group,
		Consumer: consumer,
		Streams:  []string{x.stream, ">"}, // ">" = read new messages only
		Count:    count,
		Block:    blockDuration,
	}).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, fmt.Errorf("failed to fetch new messages: %w", err)
	}

	if errors.Is(err, redis.Nil) || len(messages) == 0 {
		return []MessageWrapper[T]{}, nil
	}
	return x.parseStreamMessages(messages)
}

// parseStreamMessages parses Redis stream messages into typed MessageWrapper objects
func (x *Subscriber[T]) parseStreamMessages(streamResults []redis.XStream) ([]MessageWrapper[T], error) {
	var results []MessageWrapper[T]

	for _, streamRes := range streamResults {
		for _, msg := range streamRes.Messages {
			var payload T
			rawJSON, ok := msg.Values["data"].(string)
			if !ok {
				return nil, fmt.Errorf("payload is not a string: %v", msg.Values)
			}
			if err := json.Unmarshal([]byte(rawJSON), &payload); err != nil {
				return nil, err
			}
			results = append(results, MessageWrapper[T]{
				ID:      msg.ID,
				Payload: payload,
				Raw:     msg.Values,
			})
		}
	}
	return results, nil
}
