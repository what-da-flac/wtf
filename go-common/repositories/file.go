package repositories

import (
	"github.com/what-da-flac/wtf/openapi/domains"
)

func (x *PgRepo) InsertFile(file *domains.File) error {
	db := x.GORM()
	return db.Create(fileToDto(file)).Error
}
