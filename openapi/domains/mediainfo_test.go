package domains

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMediaInfo(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    *MediaInfo
		wantErr bool
	}{
		//{
		//	name: "new order - crystal",
		//	args: args{
		//		r: bytes.NewBufferString(`
		//{
		//  "creatingLibrary": {
		//    "name": "MediaInfoLib",
		//    "version": "25.03",
		//    "url": "https://mediaarea.net/MediaInfo"
		//  },
		//  "media": {
		//    "@ref": "01 - Crystal.flac",
		//    "track": [
		//      {
		//        "@type": "General",
		//        "AudioCount": "1",
		//        "FileExtension": "flac",
		//        "Format": "FLAC",
		//        "FileSize": "51042170",
		//        "Duration": "411.267",
		//        "OverallBitRate_Mode": "VBR",
		//        "OverallBitRate": "992877",
		//        "StreamSize": "0",
		//        "Title": "Crystal",
		//        "Album": "Get Ready",
		//        "Album_Performer": "New Order",
		//        "Part": "1",
		//        "Part_Position_Total": "1",
		//        "Track": "Crystal",
		//        "Track_Position": "1",
		//        "Track_Position_Total": "10",
		//        "Performer": "New Order",
		//        "Genre": "Alternative",
		//        "Recorded_Date": "2001",
		//        "File_Modified_Date": "2020-08-30 20:01:34 UTC",
		//        "File_Modified_Date_Local": "2020-08-30 22:01:34"
		//      },
		//      {
		//        "@type": "Audio",
		//        "Format": "FLAC",
		//        "Duration": "411.267",
		//        "BitRate_Mode": "VBR",
		//        "BitRate": "992701",
		//        "Channels": "2",
		//        "ChannelPositions": "Front: L R",
		//        "ChannelLayout": "L R",
		//        "SamplingRate": "44100",
		//        "SamplingCount": "18136860",
		//        "BitDepth": "16",
		//        "Compression_Mode": "Lossless",
		//        "StreamSize": "51033151",
		//        "Encoded_Library": "Flake#0.1",
		//        "extra": {
		//          "MD5_Unencoded": "4B3A00DB2FFD8D9183F03B3A51816569"
		//        }
		//      }
		//    ]
		//  }
		//}
		//`),
		//			},
		//			want: &MediaInfo{
		//				CreatingLibrary: CreatingLibrary{
		//					Name:    "MediaInfoLib",
		//					Version: "25.03",
		//					URL:     "https://mediaarea.net/MediaInfo",
		//				},
		//				Media: Media{
		//					Ref: "01 - Crystal.flac",
		//					Track: []Track{
		//						{
		//							Type:                  "General",
		//							AudioCount:            "1",
		//							FileExtension:         "flac",
		//							Format:                "FLAC",
		//							FileSize:              "51042170",
		//							Duration:              "411.267",
		//							OverallBitRate:        "992877",
		//							OverallBitRateMode:    "VBR",
		//							StreamSize:            "0",
		//							Title:                 "Crystal",
		//							Album:                 "Get Ready",
		//							AlbumPerformer:        "New Order",
		//							Part:                  "1",
		//							PartPositionTotal:     "1",
		//							TrackPositionTotal:    "10",
		//							TrackName:             "Crystal",
		//							TrackPosition:         "1",
		//							Performer:             "New Order",
		//							Genre:                 "Alternative",
		//							RecordedDate:          "2001",
		//							FileModifiedDate:      "2020-08-30 20:01:34 UTC",
		//							FileModifiedDateLocal: "2020-08-30 22:01:34",
		//						},
		//						{
		//							Type:             "Audio",
		//							Format:           "FLAC",
		//							Duration:         "411.267",
		//							BitRateMode:      "VBR",
		//							BitRate:          "992701",
		//							Channels:         "2",
		//							ChannelPositions: "Front: L R",
		//							ChannelLayout:    "L R",
		//							SamplingRate:     "44100",
		//							SamplingCount:    "18136860",
		//							BitDepth:         "16",
		//							CompressionMode:  "Lossless",
		//							StreamSize:       "51033151",
		//							EncodedLibrary:   "Flake#0.1",
		//						},
		//					},
		//				},
		//			},
		//			wantErr: false,
		//		},
		{
			name: "journey - whos crying now",
			args: args{
				r: bytes.NewBufferString(`
{
  "creatingLibrary": {
    "name": "MediaInfoLib",
    "version": "25.03",
    "url": "https://mediaarea.net/MediaInfo"
  },
  "media": {
    "@ref": "03 - Who's Crying Now.flac",
    "track": [
      {
        "@type": "General",
        "AudioCount": "1",
        "FileExtension": "flac",
        "Format": "FLAC",
        "FileSize": "30664257",
        "Duration": "301.827",
        "OverallBitRate_Mode": "VBR",
        "OverallBitRate": "812764",
        "StreamSize": "0",
        "Title": "Who's Crying Now",
        "Album": "Escape",
        "Album_Performer": "Journey",
        "Part": "1",
        "Part_Position_Total": "1",
        "Track": "Who's Crying Now",
        "Track_Position": "3",
        "Track_Position_Total": "10",
        "Performer": "Journey",
        "Genre": "Rock",
        "Recorded_Date": "1981",
        "File_Modified_Date": "2020-08-30 19:45:11 UTC",
        "File_Modified_Date_Local": "2020-08-30 21:45:11",
        "Comment": "Ripped by vlad198401",
        "extra": {
          "discId": "710A0B0A"
        }
      },
      {
        "@type": "Audio",
        "Format": "FLAC",
        "Duration": "301.827",
        "BitRate_Mode": "VBR",
        "BitRate": "812523",
        "Channels": "2",
        "ChannelPositions": "Front: L R",
        "ChannelLayout": "L R",
        "SamplingRate": "44100",
        "SamplingCount": "13310556",
        "BitDepth": "16",
        "Compression_Mode": "Lossless",
        "StreamSize": "30655159",
        "Encoded_Library": "reference libFLAC 1.3.3 20190804",
        "Encoded_Library_Name": "libFLAC",
        "Encoded_Library_Version": "1.3.3",
        "Encoded_Library_Date": "2019-08-04",
        "extra": {
          "MD5_Unencoded": "1409781CF0A0C9B8D517A0D793047183"
        }
      }
    ]
  }
}
`),
			},
			want: &MediaInfo{
				CreatingLibrary: CreatingLibrary{
					Name:    "MediaInfoLib",
					Version: "25.03",
					URL:     "https://mediaarea.net/MediaInfo",
				},
				Media: Media{
					Ref: "03 - Who's Crying Now.flac",
					Track: []Track{
						{
							Type:                  "General",
							AudioCount:            "1",
							FileExtension:         "flac",
							Format:                "FLAC",
							FileSize:              "30664257",
							Duration:              "301.827",
							OverallBitRateMode:    "VBR",
							OverallBitRate:        "812764",
							StreamSize:            "0",
							Title:                 "Who's Crying Now",
							Album:                 "Escape",
							AlbumPerformer:        "Journey",
							Part:                  "1",
							PartPositionTotal:     "1",
							TrackPosition:         "3",
							TrackPositionTotal:    "10",
							Performer:             "Journey",
							Genre:                 "Rock",
							RecordedDate:          "1981",
							FileModifiedDate:      "2020-08-30 19:45:11 UTC",
							FileModifiedDateLocal: "2020-08-30 21:45:11",
							TrackName:             "Who's Crying Now",
						},
						{
							Type:             "Audio",
							Format:           "FLAC",
							Duration:         "301.827",
							BitRateMode:      "VBR",
							BitRate:          "812523",
							Channels:         "2",
							ChannelPositions: "Front: L R",
							ChannelLayout:    "L R",
							SamplingRate:     "44100",
							SamplingCount:    "13310556",
							BitDepth:         "16",
							CompressionMode:  "Lossless",
							StreamSize:       "30655159",
							EncodedLibrary:   "reference libFLAC 1.3.3 20190804",
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMediaInfo(tt.args.r)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
