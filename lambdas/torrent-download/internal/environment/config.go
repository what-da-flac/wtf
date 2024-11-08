package environment

import (
	"time"

	"github.com/spf13/viper"
	"github.com/what-da-flac/wtf/go-common/env"
)

const (
	envVarBucketParsed    = "BUCKET_PARSED"
	envVarDownloadsBucket = "BUCKET_DOWNLOADS"
	envVarTimeout         = "TIMEOUT"
)

type Config struct {
	*env.Config
	BucketDownloads string
	BucketParsed    string
	Timeout         time.Duration
}

func New() *Config {
	globalConfig := env.New()
	return &Config{
		Config:          globalConfig,
		BucketDownloads: viper.GetString(envVarDownloadsBucket),
		BucketParsed:    viper.GetString(envVarBucketParsed),
		Timeout:         viper.GetDuration(envVarTimeout),
	}
}
