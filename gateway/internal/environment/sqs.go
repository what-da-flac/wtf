package environment

import "github.com/spf13/viper"

const (
	envVarTorrentMetadataUrl = "QUEUE_TORRENT_METADATA_URL"
	envVarTorrentParsedUrl   = "QUEUE_TORRENT_PARSED_URL"
)

type SQS struct {
	TorrentMetadataUrl string
	TorrentParsedUrl   string
}

func newSQS() SQS {
	return SQS{
		TorrentMetadataUrl: viper.GetString(envVarTorrentMetadataUrl),
		TorrentParsedUrl:   viper.GetString(envVarTorrentParsedUrl),
	}
}
