package cmd

import (
	"time"

	"github.com/what-da-flac/wtf/go-common/ifaces"
	"github.com/what-da-flac/wtf/go-common/loggers"
	"github.com/what-da-flac/wtf/go-common/rabbits"
	"github.com/what-da-flac/wtf/services/gateway/internal/environment"

	"github.com/spf13/cobra"
)

var rabbitListenerCmd = &cobra.Command{
	Use:   "rabbit-listener",
	Short: "Listens to rabbitmq messages",
	Run: func(cmd *cobra.Command, args []string) {
		rabbitListener()
	},
}

func init() {
	rootCmd.AddCommand(rabbitListenerCmd)
}

func rabbitListener() {
	logger := loggers.MustNewDevelopmentLogger()
	config := environment.New()
	l := rabbits.NewListener(logger, "test-queue", config.RabbitMQ.URL, time.Second)
	l.ListenAsync(func(msg []byte) (ack ifaces.AckType, err error) {
		logger.Info("received:", string(msg))
		return ifaces.MessageAcknowledge, nil
	})
	select {}
}
