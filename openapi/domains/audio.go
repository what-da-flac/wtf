package domains

import (
	"math"
	"strconv"
	"time"
)

type Audio struct {
	Album           string
	BitDepth        int
	CompressionMode string
	Duration        time.Duration
	FileExtension   string
	Format          string
	Genre           string
	Performer       string
	RecordedDate    int
	SamplingRate    int
	Title           string
	TrackNumber     int
	TotalTrackCount int
}

func NewAudio(info MediaInfo) Audio {
	var r Audio
	if track := info.Audio(); track != nil {
		if val, err := strconv.Atoi(track.BitDepth); err == nil {
			r.BitDepth = val
		}
		r.CompressionMode = track.CompressionMode
		if val, err := strconv.ParseFloat(track.Duration, 64); err == nil {
			r.Duration = time.Second * time.Duration(math.Floor(val))
		}
		r.Format = track.Format
		if val, err := strconv.Atoi(track.SamplingRate); err == nil {
			r.SamplingRate = val
		}
	}
	if track := info.General(); track != nil {
		r.Album = track.Album
		r.FileExtension = track.FileExtension
		r.Genre = track.Genre
		r.Performer = track.Performer
		if val, err := strconv.Atoi(track.RecordedDate); err == nil {
			r.RecordedDate = val
		}
		r.Title = track.Title
		if val, err := strconv.Atoi(track.TrackPosition); err == nil {
			r.TrackNumber = val
		}
		if val, err := strconv.Atoi(track.TrackPositionTotal); err == nil {
			r.TotalTrackCount = val
		}
	}

	return r
}
