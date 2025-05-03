package main

import (
	"encoding/json"
	"fmt"
	"log"

	"go.uber.org/zap"

	"golang.org/x/net/context"

	"github.com/what-da-flac/wtf/go-common/brokers"
	"github.com/what-da-flac/wtf/go-common/env"
	"github.com/what-da-flac/wtf/openapi/gen/golang"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	logger, err := zap.NewProductionConfig().Build()
	if err != nil {
		return err
	}
	return serve(logger)
}

func serve(zl *zap.Logger) error {
	ctx := context.Background()
	config := env.New()
	// TODO: set redis connection from environment variables
	_ = config
	client := brokers.NewClient()
	queueName := string(golang.QueueNameMediainfo)
	subscriber, err := brokers.NewSubscriber[golang.MediaInfoInput](client, queueName, "media-info")
	if err != nil {
		return err
	}
	processMessageFn := func(msg golang.MediaInfoInput) (ack bool, err error) {
		// TODO: read mediainfo
		// TODO: if audio resolution is below msg.MinBitrate, acknowledge with error
		// TODO: write final audio file to db
		data, err := json.Marshal(msg)
		if err != nil {
			return false, err
		}
		fmt.Println("received message:", string(data))
		return true, nil
	}
	errFn := func(err error) {
		log.Fatal(err)
	}
	zl.Sugar().Infoln("starting subscriber:", queueName)
	subscriber.Listen(ctx, processMessageFn, errFn)
	return nil
}
