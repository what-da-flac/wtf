package cmd

import (
	"fmt"
	"github.com/what-da-flac/wtf/go-common/loggers"
	"github.com/what-da-flac/wtf/go-common/rabbits"
	"github.com/what-da-flac/wtf/services/gateway/internal/environment"
	"time"

	"github.com/spf13/cobra"
)

var rabbitListenerCmd = &cobra.Command{
	Use:   "rabbit-listener",
	Short: "Listens to rabbitmq messages",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("not implemented")
	},
}

func init() {
	rootCmd.AddCommand(rabbitListenerCmd)
}

func rabbitListener() {
	logger := loggers.MustNewDevelopmentLogger()
	config := environment.New()
	rabbits.NewListener(logger, "test-queue", config.RabbitMQ.URL, time.Second)
	// TODO: continue with listener and publisher
}
