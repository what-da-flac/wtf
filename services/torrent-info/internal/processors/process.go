package processors

import (
	"github.com/what-da-flac/wtf/go-common/env"
	"github.com/what-da-flac/wtf/go-common/ifaces"
	"github.com/what-da-flac/wtf/openapi/models"
)

func Process(logger ifaces.Logger, config *env.Config, torrent *models.Torrent) error {
	logger.Infof("torrent name: %s file count: %d", torrent.Name, len(torrent.Files))
	logger.Info("TODO: store in database")
	return nil
}
