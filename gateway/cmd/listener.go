package cmd

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/what-da-flac/wtf/gateway/internal/environment"
	"github.com/what-da-flac/wtf/gateway/internal/interfaces"
	"github.com/what-da-flac/wtf/gateway/internal/listeners"
	"github.com/what-da-flac/wtf/gateway/internal/repositories"
	"github.com/what-da-flac/wtf/gateway/internal/senders"
	"github.com/what-da-flac/wtf/go-common/amazon"
	"github.com/what-da-flac/wtf/go-common/identifiers"
	"github.com/what-da-flac/wtf/go-common/pgpq"
	"github.com/what-da-flac/wtf/go-common/timers"
	"go.uber.org/zap"
)

var listenCmd = &cobra.Command{
	Use:   "listener",
	Short: "Listen one or more sqs queues for messages",
	RunE: func(cmd *cobra.Command, args []string) error {
		zl, err := zap.NewProductionConfig().Build()
		if err != nil {
			return err
		}
		config := environment.New()
		return Listen(config, zl.Sugar())
	},
}

func init() {
	rootCmd.AddCommand(listenCmd)
}

func Listen(config *environment.Config, logger interfaces.Logger) error {
	var messageListeners []interfaces.MessageListener
	awsSession := amazon.NewAWSSessionFromEnvironment()
	if err := awsSession.Build(); err != nil {
		return err
	}
	connStr := config.DB.URL
	db, err := pgpq.New(connStr)
	if err != nil {
		return err
	}
	defer func() { _ = db.Close() }()
	logger.Info("connected to db")
	repository, err := repositories.NewPG(db, connStr)
	if err != nil {
		return err
	}
	sess := awsSession.Session()
	identifier := identifiers.NewIdentifier()
	sender := senders.NewMessageSender(sess, logger, identifier)
	timer := timers.New()
	jobs := listeners.NewJobs(sess, config, logger, repository, sender, timer, identifier)
	if err := jobs.Build(); err != nil {
		return err
	}
	jobItems := jobs.Map()
	for k, v := range jobItems {
		name := k
		messageListeners = append(messageListeners,
			listeners.NewSQSListener(
				sess, v.Fn, name,
				v.ListenerUri, logger,
				v.VisibilityTimeout,
				v.WaitTime,
				v.MaxNumberOfMessages,
			),
		)
	}
	// Use a WaitGroup to manage concurrent listeners
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Capture OS signals for graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	for _, listener := range messageListeners {
		wg.Add(1)
		go func(l interfaces.MessageListener) {
			defer wg.Done()
			// start polling messages
			if err := l.Poll(ctx); err != nil {
				logger.Errorf("error while polling messages: %v", err)
			}
		}(listener)
	}

	// Wait for termination signal
	select {
	case sig := <-signals:
		logger.Infof("Received signal: %s, initiating shutdown...", sig)
		cancel()
	case <-ctx.Done():
		logger.Info("context done, shutting down...")
	}

	// Wait for all listeners to finish
	wg.Wait()
	logger.Info("all listeners shut down gracefully.")
	return nil
}
