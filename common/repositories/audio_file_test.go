package repositories

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPgRepo_UpdateAudioFile(t *testing.T) {
	repo, err := InitializeConn()
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
