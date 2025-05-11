package repositories

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestPgRepo_UpdateAudioFile(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)
	repo := PgRepo{
		_db: db,
	}
	err = db.AutoMigrate(&AudioFileDto{})
	require.NoError(t, err)

	dto := AudioFileDto{
		Id:     "abc-123",
		Title:  "Original Title",
		Length: 123456,
		Genre:  "Rock",
	}
	values := map[string]any{
		"genre":  dto.Genre,
		"length": dto.Length,
		"title":  "New Title",
	}
	err = repo.UpdateAudioFile(dto.Id, values)
	assert.NoError(t, err)
}
