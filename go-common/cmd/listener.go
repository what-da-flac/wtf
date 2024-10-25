package cmd

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/what-da-flac/wtf/go-common/environment"
	"github.com/what-da-flac/wtf/go-common/ifaces"
	"github.com/what-da-flac/wtf/go-common/rabbit"
	"go.uber.org/zap"
)

var listenCmd = &cobra.Command{
	Use:   "listener",
	Short: "Listens to rabbitmq messages",
	RunE: func(cmd *cobra.Command, args []string) error {
		return Listener()
	},
}

func init() {
	rootCmd.AddCommand(listenCmd)
}

func Listener() error {
	lg, err := zap.NewProduction()
	if err != nil {
		return err
	}
	logger := lg.Sugar()
	rmq := environment.New().RabbitMQ
	name := viper.GetString("QUEUE_NAME")
	listener := rabbit.NewListener().
		WithConnection(rmq.Protocol, rmq.Username, rmq.Password, rmq.Hostname, rmq.Port).
		WithName(name).
		WithAckErrorHandler(func(msg []byte, err error) {
			logger.Error(err)
		}).WithHandler(listener(logger))
	if err := listener.Build(); err != nil {
		return err
	}
	logger.Infof("starting listener")
	listener.ListenAsync()
	interval := time.Minute
	timer := time.After(interval)
	for {
		<-timer
		logger.Info("timer expired")
		timer = time.After(interval)
	}
}

func listener(lg *zap.SugaredLogger) func(msg []byte) (ifaces.AckType, error) {
	return func(msg []byte) (ifaces.AckType, error) {
		payload := string(msg)
		lg.Infof("incoming message: %s", payload)
		switch payload {
		case "":
			return ifaces.MessageReject, nil
		case "test":
			return ifaces.MessageRequeue, nil
		default:
			return ifaces.MessageAcknowledge, nil
		}
	}
}
