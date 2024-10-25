package cmd

import (
	"context"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/what-da-flac/wtf/go-common/environment"
	"github.com/what-da-flac/wtf/go-common/rabbit"
)

var publisherCmd = &cobra.Command{
	Use:   "publisher",
	Short: "Publishes a message to rabbitmq broker",
	RunE: func(cmd *cobra.Command, args []string) error {
		return Publisher(args)
	},
}

func init() {
	cmd := publisherCmd
	rootCmd.AddCommand(cmd)
}

func Publisher(args []string) error {
	var (
		count int = 1
		name  string
	)
	ctx := context.TODO()
	name = "hello"
	if len(args) > 0 {
		name = args[0]
	}
	if len(args) > 1 {
		if v, err := strconv.Atoi(args[1]); err == nil && v > 0 {
			count = v
		}
	}
	config := environment.New()
	rmq := config.RabbitMQ
	queue := viper.GetString("QUEUE_NAME")
	publisher := rabbit.NewPublisher().
		WithConnection(rmq.Protocol, rmq.Username, rmq.Password, rmq.Hostname, rmq.Port).
		WithName(queue)
	if err := publisher.Build(); err != nil {
		return err
	}
	for i := 0; i < count; i++ {
		if err := publisher.Publish(ctx, []byte(name)); err != nil {
			return err
		}
	}
	return nil
}
