package listeners

import (
	"context"
	"encoding/json"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/what-da-flac/wtf/gateway/internal/domain/torrent"
	"github.com/what-da-flac/wtf/gateway/internal/environment"
	"github.com/what-da-flac/wtf/gateway/internal/interfaces"
	"github.com/what-da-flac/wtf/openapi/models"
)

// Job represents a listener job.
type Job struct {
	Fn                  interfaces.MessageReceiverFn
	MaxNumberOfMessages int
	ListenerUri         string
	VisibilityTimeout   time.Duration
	WaitTime            time.Duration
}

// Jobs encapsulates multiple Jobs.
type Jobs struct {
	config     *environment.Config
	identifier interfaces.Identifier
	logger     interfaces.Logger
	sess       *session.Session
	sender     interfaces.MessageSender

	// interfaces jobs need to implement
	torrentUpdate interface {
		Save(ctx context.Context, torrent *models.Torrent) error
	}
}

func NewJobs(sess *session.Session, config *environment.Config, logger interfaces.Logger,
	repository interfaces.Repository, sender interfaces.MessageSender,
	timer interfaces.Timer, identifier interfaces.Identifier) *Jobs {
	return &Jobs{
		config:        config,
		identifier:    identifier,
		logger:        logger,
		sender:        sender,
		sess:          sess,
		torrentUpdate: torrent.NewUpdate(repository, timer, identifier),
	}
}

func (x *Jobs) Build() error {
	return nil
}

// Map returns a map of listener jobs.
//
//nolint:mnd,unused
func (x *Jobs) Map() map[string]*Job {
	const (
		defaultVisibilityTimeout = 10 * time.Second
		defaultWaitTime          = 5 * time.Second
	)
	return map[string]*Job{
		// receives parsed torrent
		"TorrentParsed": {
			Fn: func(body string) error {
				payload := &models.Torrent{}
				if err := json.Unmarshal([]byte(body), payload); err != nil {
					return err
				}
				x.logger.Infof("received torrent id: %s name: %s", payload.Id, payload.Name)
				ctx := context.Background()
				// TODO: report to sentry on any error below
				err := x.torrentUpdate.Save(ctx, payload)
				if err != nil {
					x.logger.Errorf("failed to save torrent id: %s name: %s", payload.Id, payload.Name)
				}
				return err
			},
			ListenerUri:         x.config.SQS.TorrentParsedUrl,
			MaxNumberOfMessages: 1,
			VisibilityTimeout:   time.Minute,
			WaitTime:            time.Second * 20,
		},
	}
}