package cmd

import (
	"github.com/what-da-flac/wtf/go-common/loggers"
	"github.com/what-da-flac/wtf/go-common/rabbits"
	"github.com/what-da-flac/wtf/services/gateway/internal/environment"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var rabbitPublisherCmd = &cobra.Command{
	Use:   "rabbit-publisher",
	Short: "Publishes rabbitmq messages",
	Run: func(cmd *cobra.Command, args []string) {
		if err := rabbitPublisher(); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(rabbitPublisherCmd)
}

func rabbitPublisher() error {
	data := `test message`
	if args := os.Args; len(args) != 0 {
		data = strings.Join(args, " ")
	}
	logger := loggers.MustNewDevelopmentLogger()
	config := environment.New()
	publisher := rabbits.NewPublisher(logger, "test-queue", config.RabbitMQ.URL)
	if err := publisher.Build(); err != nil {
		return err
	}
	return publisher.Publish([]byte(data))
}
