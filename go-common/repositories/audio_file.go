package repositories

import (
	"github.com/what-da-flac/wtf/openapi/domains"
)

func (x *PgRepo) InsertAudioFile(file *domains.AudioFile) error {
	db := x.GORM()
	return db.Create(fileToDto(file)).Error
}
