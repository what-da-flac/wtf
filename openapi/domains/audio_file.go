package domains

import (
	"github.com/what-da-flac/wtf/openapi/gen/golang"
)

func NewAudioFile(audio *golang.Audio, file *golang.File) golang.AudioFile {
	return golang.AudioFile{
		Album:           audio.Album,
		BitDepth:        audio.BitDepth,
		CompressionMode: audio.CompressionMode,
		Duration:        audio.Duration,
		FileExtension:   audio.FileExtension,
		Format:          audio.Format,
		Genre:           audio.Genre,
		Performer:       audio.Performer,
		RecordedDate:    audio.RecordedDate,
		SamplingRate:    audio.SamplingRate,
		Title:           audio.Title,
		TrackNumber:     audio.TrackNumber,
		TotalTrackCount: audio.TotalTrackCount,
		Id:              file.Id,
		Filename:        file.Filename,
		Created:         file.Created,
		Length:          file.Length,
		ContentType:     file.ContentType,
		Status:          file.Status,
	}
}
