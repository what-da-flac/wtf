package repositories

import (
	"github.com/what-da-flac/wtf/openapi/gen/golang"
)

func (x *PgRepo) InsertAudioFile(file *golang.AudioFile) error {
	db := x.GORM()
	dto := fileToDto(file)
	return db.Create(dto).Error
}

func (x *PgRepo) SelectAudioFile(id string) (*golang.AudioFile, error) {
	db := x.GORM()
	dto := &AudioFileDto{
		Id: id,
	}
	if err := db.Find(dto).Error; err != nil {
		return nil, err
	}
	return dto.toFile(), nil
}

func (x *PgRepo) UpdateAudioFile(id string, values map[string]any) error {
	db := x.GORM()
	dto := &AudioFileDto{
		Id: id,
	}
	return db.Model(dto).Updates(values).Error
}
