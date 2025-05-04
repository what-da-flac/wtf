package repositories

import (
	"github.com/what-da-flac/wtf/openapi/gen/golang"
)

func (x *PgRepo) InsertAudioFile(file *golang.AudioFile) error {
	db := x.GORM()
	return db.Create(fileToDto(file)).Error
}
