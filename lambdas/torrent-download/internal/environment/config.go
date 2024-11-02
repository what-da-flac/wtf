package environment

import (
	"time"

	"github.com/spf13/viper"
	"github.com/what-da-flac/wtf/go-common/environment"
)

const (
	envVarSQSTorrentParsedURL   = "SQS_TORRENT_PARSED_URL"
	envVarS3TorrentParsedBucket = "S3_TORRENT_PARSED_BUCKET"
	envVarTimeout               = "TIMEOUT"
	envVarTorrentBucket         = "BUCKET_TORRENT"
	envVarDownloadsBucket       = "BUCKET_DOWNLOADS"
)

type Config struct {
	*environment.Config
	BucketDownloads       string
	BucketTorrent         string
	SQSTorrentParsedURL   string
	S3TorrentParsedBucket string
	Timeout               time.Duration
}

func New() *Config {
	globalConfig := environment.New()
	return &Config{
		Config:                globalConfig,
		BucketDownloads:       viper.GetString(envVarDownloadsBucket),
		BucketTorrent:         viper.GetString(envVarTorrentBucket),
		SQSTorrentParsedURL:   viper.GetString(envVarSQSTorrentParsedURL),
		S3TorrentParsedBucket: viper.GetString(envVarS3TorrentParsedBucket),
		Timeout:               viper.GetDuration(envVarTimeout),
	}
}
