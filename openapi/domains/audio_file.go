package domains

import "time"

type AudioFile struct {
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
	Id              string
	Filename        string
	Created         time.Time
	Length          int64
	ContentType     string
	Status          string
}

func NewAudioFile(audio *Audio, file *File) AudioFile {
	return AudioFile{
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
