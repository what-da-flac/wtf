package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/what-da-flac/wtf/common/identifiers"
)

func TestCmdFFMpegSetTags(t *testing.T) {
	identifier := identifiers.NewIdentifier()
	filename := "/Users/mau/Downloads/stuff/audio/Arrival/01 Mosquito Bite.m4a"
	values := map[string]string{
		"composer": "Mauricio Leyzaola",
		"artist":   "Mike Stern",
		"title":    "Photograph",
	}
	err := CmdFFMpegSetTags(identifier, filename, values)
	assert.NoError(t, err)
}
