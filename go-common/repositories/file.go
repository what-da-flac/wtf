package repositories

import (
	"time"

	"github.com/what-da-flac/wtf/openapi/domains"
)

type FileDto struct {
	Id          string
	Filename    string
	Created     time.Time
	Length      int64
	ContentType string
	Status      string
}

func (x *FileDto) TableName() string { return "files" }

func (x *PgRepo) InsertFile(file *domains.File) error {
	db := x.GORM()
	return db.Create(file).Error
}
