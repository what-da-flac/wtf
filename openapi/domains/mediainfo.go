package domains

import (
	"encoding/json"
	"io"
)

type MediaInfo struct {
	CreatingLibrary CreatingLibrary `json:"creatingLibrary"`
	Media           Media           `json:"media"`
}

type CreatingLibrary struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	URL     string `json:"url"`
}

type Media struct {
	Ref   string  `json:"@ref"`
	Track []Track `json:"track"`
}

type Track struct {
	Album                 string `json:"Album,omitempty"`
	AlbumPerformer        string `json:"Album_Performer,omitempty"`
	AudioCount            string `json:"AudioCount,omitempty"`
	BitDepth              string `json:"BitDepth,omitempty"`
	BitRate               string `json:"BitRate,omitempty"`
	BitRateMode           string `json:"BitRate_Mode,omitempty"`
	ChannelLayout         string `json:"ChannelLayout,omitempty"`
	ChannelPositions      string `json:"ChannelPositions,omitempty"`
	Channels              string `json:"Channels,omitempty"`
	CompressionMode       string `json:"Compression_Mode,omitempty"`
	Duration              string `json:"Duration"`
	EncodedLibrary        string `json:"Encoded_Library,omitempty"`
	FileExtension         string `json:"FileExtension,omitempty"`
	FileModifiedDate      string `json:"File_Modified_Date,omitempty"`
	FileModifiedDateLocal string `json:"File_Modified_Date_Local,omitempty"`
	FileSize              string `json:"FileSize,omitempty"`
	Format                string `json:"Format"`
	Genre                 string `json:"Genre,omitempty"`
	OverallBitRate        string `json:"OverallBitRate,omitempty"`
	OverallBitRateMode    string `json:"OverallBitRate_Mode,omitempty"`
	Part                  string `json:"Part,omitempty"`
	PartPositionTotal     string `json:"Part_Position_Total,omitempty"`
	Performer             string `json:"Performer,omitempty"`
	RecordedDate          string `json:"Recorded_Date,omitempty"`
	SamplingCount         string `json:"SamplingCount,omitempty"`
	SamplingRate          string `json:"SamplingRate,omitempty"`
	StreamSize            string `json:"StreamSize,omitempty"`
	Title                 string `json:"Title,omitempty"`
	TrackName             string `json:"Track,omitempty"`
	TrackPosition         string `json:"Track_Position,omitempty"`
	TrackPositionTotal    string `json:"Track_Position_Total,omitempty"`
	Type                  string `json:"@type"`
}

func NewMediaInfo(r io.Reader) (*MediaInfo, error) {
	x := &MediaInfo{}
	if err := json.NewDecoder(r).Decode(x); err != nil {
		return nil, err
	}
	return x, nil
}

func (x MediaInfo) Audio() *Track {
	for _, track := range x.Media.Track {
		if track.Type == "Audio" {
			return &track
		}
	}
	return nil
}

func (x MediaInfo) General() *Track {
	for _, track := range x.Media.Track {
		if track.Type == "General" {
			return &track
		}
	}
	return nil
}
