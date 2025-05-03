package domains

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/what-da-flac/wtf/openapi/gen/golang"
)

func TestNewAudio(t *testing.T) {
	type args struct {
		info *MediaInfo
	}
	tests := []struct {
		name string
		args args
		want golang.Audio
	}{
		{
			name: "new order - crystal",
			args: args{
				info: &MediaInfo{
					CreatingLibrary: CreatingLibrary{
						Name:    "MediaInfoLib",
						Version: "25.03",
						URL:     "https://mediaarea.net/MediaInfo",
					},
					Media: Media{
						Ref: "01 - Crystal.flac",
						Track: []Track{
							{
								Type:                  "General",
								AudioCount:            "1",
								FileExtension:         "flac",
								Format:                "FLAC",
								FileSize:              "51042170",
								Duration:              "411.267",
								OverallBitRate:        "992877",
								OverallBitRateMode:    "VBR",
								StreamSize:            "0",
								Title:                 "Crystal",
								Album:                 "Get Ready",
								AlbumPerformer:        "New Order",
								Part:                  "1",
								PartPositionTotal:     "1",
								TrackPositionTotal:    "10",
								TrackName:             "Crystal",
								TrackPosition:         "1",
								Performer:             "New Order",
								Genre:                 "Alternative",
								RecordedDate:          "2001",
								FileModifiedDate:      "2020-08-30 20:01:34 UTC",
								FileModifiedDateLocal: "2020-08-30 22:01:34",
							},
							{
								Type:             "Audio",
								Format:           "FLAC",
								Duration:         "411.267",
								BitRateMode:      "VBR",
								BitRate:          "992701",
								Channels:         "2",
								ChannelPositions: "Front: L R",
								ChannelLayout:    "L R",
								SamplingRate:     "44100",
								SamplingCount:    "18136860",
								BitDepth:         "16",
								CompressionMode:  "Lossless",
								StreamSize:       "51033151",
								EncodedLibrary:   "Flake#0.1",
							},
						},
					},
				},
			},
			want: golang.Audio{
				Album:           "Get Ready",
				BitDepth:        16,
				CompressionMode: "Lossless",
				DurationSeconds: 411,
				FileExtension:   "flac",
				Format:          "FLAC",
				Genre:           "Alternative",
				Performer:       "New Order",
				RecordedDate:    2001,
				SamplingRate:    44100,
				Title:           "Crystal",
				TrackNumber:     1,
				TotalTrackCount: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewAudio(tt.args.info), "NewAudio(%v)", tt.args.info)
		})
	}
}
