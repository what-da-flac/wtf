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
	Type                  string `json:"@type"`
	AudioCount            string `json:"AudioCount,omitempty"`
	FileExtension         string `json:"FileExtension,omitempty"`
	Format                string `json:"Format"`
	FileSize              string `json:"FileSize,omitempty"`
	Duration              string `json:"Duration"`
	OverallBitRateMode    string `json:"OverallBitRate_Mode,omitempty"`
	OverallBitRate        string `json:"OverallBitRate,omitempty"`
	StreamSize            string `json:"StreamSize,omitempty"`
	Title                 string `json:"Title,omitempty"`
	Album                 string `json:"Album,omitempty"`
	AlbumPerformer        string `json:"Album_Performer,omitempty"`
	Part                  string `json:"Part,omitempty"`
	PartPositionTotal     string `json:"Part_Position_Total,omitempty"`
	TrackName             string `json:"Track,omitempty"`
	TrackPosition         string `json:"Track_Position,omitempty"`
	TrackPositionTotal    string `json:"Track_Position_Total,omitempty"`
	Performer             string `json:"Performer,omitempty"`
	Genre                 string `json:"Genre,omitempty"`
	RecordedDate          string `json:"Recorded_Date,omitempty"`
	FileModifiedDate      string `json:"File_Modified_Date,omitempty"`
	FileModifiedDateLocal string `json:"File_Modified_Date_Local,omitempty"`
	BitRateMode           string `json:"BitRate_Mode,omitempty"`
	BitRate               string `json:"BitRate,omitempty"`
	Channels              string `json:"Channels,omitempty"`
	ChannelPositions      string `json:"ChannelPositions,omitempty"`
	ChannelLayout         string `json:"ChannelLayout,omitempty"`
	SamplingRate          string `json:"SamplingRate,omitempty"`
	SamplingCount         string `json:"SamplingCount,omitempty"`
	BitDepth              string `json:"BitDepth,omitempty"`
	CompressionMode       string `json:"Compression_Mode,omitempty"`
	EncodedLibrary        string `json:"Encoded_Library,omitempty"`
}

func NewMediaInfo(r io.Reader) MediaInfo {
	x := MediaInfo{}
	if err := json.NewDecoder(r).Decode(&x); err != nil {
		return MediaInfo{}
	}
	return x
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
