package domains

import (
	"bytes"
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
			name: "amy winehose - rehab.m4a",
			args: args{
				info: func() *MediaInfo {
					payload := `{
  "creatingLibrary": {
    "name": "MediaInfoLib",
    "version": "25.07",
    "url": "https://mediaarea.net/MediaInfo"
  },
  "media": {
    "@ref": "/Users/mau/Downloads/stuff/audio/m4a/Amy Winehouse/Back to Black/01 Rehab.m4a",
    "track": [
      {
        "@type": "General",
        "AudioCount": "1",
        "ImageCount": "1",
        "FileExtension": "m4a",
        "Format": "MPEG-4",
        "Format_Profile": "Apple audio with iTunes info",
        "CodecID": "M4A ",
        "CodecID_Compatible": "M4A /mp42/isom",
        "FileSize": "7602011",
        "Duration": "214.808",
        "OverallBitRate_Mode": "VBR",
        "OverallBitRate": "283118",
        "StreamSize": "526235",
        "HeaderSize": "596807",
        "DataSize": "7005204",
        "FooterSize": "0",
        "IsStreamable": "Yes",
        "Title": "Rehab",
        "Album": "Back to Black",
        "Album_Sort": "Back to Black",
        "Album_Performer": "Amy Winehouse",
        "Part_Position": "1",
        "Part_Position_Total": "1",
        "Track": "Rehab",
        "Track_Position": "1",
        "Track_Position_Total": "11",
        "Performer": "Amy Winehouse",
        "Performer_Sort": "Amy Winehouse",
        "Composer": "Amy Winehouse",
        "Genre": "R&B/Soul",
        "ContentType": "Music",
        "Recorded_Date": "2006-10-23 07:00:00 UTC",
        "Encoded_Date": "2031-09-15 02:32:49 UTC",
        "Tagged_Date": "2019-09-14 11:57:11 UTC",
        "File_Modified_Date": "2020-08-30 18:59:40 UTC",
        "File_Modified_Date_Local": "2020-08-30 20:59:40",
        "Cover": "Yes",
        "Cover_Type": "Cover",
        "extra": {
          "AppleStoreCatalogID": "1011661392",
          "AlbumTitleID": "13125609",
          "cmID": "13125609",
          "PlayListID": "1011661390",
          "GenreID": "15",
          "PurchaseDate": "2018-02-26 04:02:00",
          "Title_Sort": "Rehab"
        }
      },
      {
        "@type": "Audio",
        "StreamOrder": "0",
        "ID": "1",
        "Format": "AAC",
        "Format_AdditionalFeatures": "LC",
        "CodecID": "mp4a-40-2",
        "Duration": "214.808",
        "BitRate_Mode": "VBR",
        "BitRate": "256000",
        "BitRate_Maximum": "425648",
        "Channels": "2",
        "ChannelPositions": "Front: L R",
        "ChannelLayout": "L R",
        "SamplesPerFrame": "1024",
        "SamplingRate": "44100",
        "SamplingCount": "9473033",
        "FrameRate": "43.066",
        "FrameCount": "9251",
        "Compression_Mode": "Lossy",
        "StreamSize": "7005196",
        "Language": "en",
        "Encoded_Date": "2031-09-15 02:32:49 UTC",
        "Tagged_Date": "2019-09-14 11:57:11 UTC",
        "extra": {
          "linf": {
            "@dt": "binary.base64",
            "#value": "AQ=="
          }
        }
      },
      {
        "@type": "Image",
        "Type": "Cover",
        "Format": "JPEG",
        "MuxingMode": "moov-meta-covr",
        "Width": "600",
        "Height": "600",
        "ColorSpace": "YUV",
        "ChromaSubsampling": "4:4:4",
        "BitDepth": "8",
        "Compression_Mode": "Lossy",
        "StreamSize": "70580"
      }
    ]
  }
}
`
					mediaInfo, err := NewMediaInfo(bytes.NewBufferString(payload))
					if err != nil {
						t.Fatal(err)
					}
					return mediaInfo
				}(),
			},
			want: golang.Audio{
				Album:           "Back to Black",
				BitDepth:        0,
				BitRate:         256000,
				CompressionMode: "Lossy",
				DurationSeconds: 214,
				FileExtension:   "m4a",
				Format:          "AAC",
				Genre:           "R&B/Soul",
				Performer:       "Amy Winehouse",
				RecordedDate:    2006,
				Title:           "Rehab",
				TotalTrackCount: 11,
				TrackNumber:     1,
			},
		},
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
				BitRate:         992701,
				CompressionMode: "Lossless",
				DurationSeconds: 411,
				FileExtension:   "flac",
				Format:          "FLAC",
				Genre:           "Alternative",
				Performer:       "New Order",
				RecordedDate:    2001,
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
