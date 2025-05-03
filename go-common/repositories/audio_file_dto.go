package repositories

import (
	"time"

	"github.com/jinzhu/copier"
	"github.com/what-da-flac/wtf/openapi/gen/golang"
)

type AudioFileDto struct {
	Id              string
	Filename        string
	Created         time.Time
	Length          int64
	ContentType     string
	Status          string
	Album           string
	BitDepth        int
	CompressionMode string
	DurationSeconds int
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

func (x *AudioFileDto) TableName() string { return "audio_files" }

func fileToDto(file *golang.AudioFile) *AudioFileDto {
	res := &AudioFileDto{}
	if err := copier.Copy(res, file); err != nil {
		return nil
	}
	return res
}

func (x *AudioFileDto) toFile() *golang.File {
	res := &golang.File{}
	if err := copier.Copy(res, x); err != nil {
		return nil
	}
	return res
}
