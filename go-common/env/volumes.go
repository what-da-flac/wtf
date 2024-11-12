package env

import "github.com/spf13/viper"

type Volumes struct {
	Downloads Names
	Media     Names
}

func newVolumes() Volumes {
	return Volumes{
		Downloads: Names(viper.GetString("VOLUME_DOWNLOADS_PATH")),
		Media:     Names(viper.GetString("VOLUME_MEDIA_PATH")),
	}
}
