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
	if err := db.Where("id = ?", id).First(dto).Error; err != nil {
		return nil, err
	}
	return dto.toAudioFile(), nil
}

func (x *PgRepo) UpdateAudioFile(id string, values map[string]any) error {
	db := x.GORM()
	dto := &AudioFileDto{
		Id: id,
	}
	return db.Model(dto).Updates(values).Error
}

func (x *PgRepo) FindByHash(hash string) (*golang.AudioFile, error) {
	db := x.GORM()
	dto := &AudioFileDto{}
	if err := db.Where("src_hash = ?", hash).First(dto).Error; err != nil {
		return nil, err
	}
	return dto.toAudioFile(), nil
}
