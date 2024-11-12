package env

import (
	"time"

	"github.com/spf13/viper"
)

type Downloads struct {
	Timeout time.Duration
}

func newDownloads() Downloads {
	return Downloads{
		Timeout: viper.GetDuration("TIMEOUT_TORRENT_DOWNLOAD"),
	}
}
