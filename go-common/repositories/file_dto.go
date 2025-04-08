package repositories

import (
	"time"

	"github.com/jinzhu/copier"

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

func fileToDto(file *domains.File) *FileDto {
	res := &FileDto{}
	if err := copier.Copy(res, file); err != nil {
		return nil
	}
	return res
}

func (x *FileDto) toFile() *domains.File {
	res := &domains.File{}
	if err := copier.Copy(res, x); err != nil {
		return nil
	}
	return res
}
