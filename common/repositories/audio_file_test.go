package repositories

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPgRepo_UpdateAudioFile(t *testing.T) {
	// this unit test doesn't really check anything, just gets the output of generated sql command
	db, _, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	repo, err := NewPgRepo(db, "postgres://", true)
	assert.NoError(t, err)

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
