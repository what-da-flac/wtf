package environment

import (
	"github.com/spf13/viper"
	"github.com/what-da-flac/wtf/go-common/environment"
)

const (
	envVarSQSTorrentParsedURL   = "SQS_TORRENT_PARSED_URL"
	envVarS3TorrentParsedBucket = "S3_TORRENT_PARSED_BUCKET"
)

type Config struct {
	*environment.Config
	SQSTorrentParsedURL   string
	S3TorrentParsedBucket string
}

func New() *Config {
	globalConfig := environment.New()
	return &Config{
		Config:                globalConfig,
		SQSTorrentParsedURL:   viper.GetString(envVarSQSTorrentParsedURL),
		S3TorrentParsedBucket: viper.GetString(envVarS3TorrentParsedBucket),
	}
}
