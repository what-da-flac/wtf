package processors

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/what-da-flac/wtf/go-common/env"
	"github.com/what-da-flac/wtf/go-common/ifaces"
	"github.com/what-da-flac/wtf/openapi/models"
)

func Process(publisher ifaces.Publisher, logger ifaces.Logger,
	sess *session.Session,
	config *env.Config, torrent *models.Torrent) error {
	return nil
}
