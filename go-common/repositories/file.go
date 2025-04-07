package repositories

import (
	"time"

	"github.com/what-da-flac/wtf/openapi/gen/golang"
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

func (x *PgRepo) InsertFile(file *golang.File) error {
	db := x.GORM()
	return db.Create(file).Error
}
