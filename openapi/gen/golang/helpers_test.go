package golang

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/what-da-flac/wtf/common/helpers"
)

func TestPatchV1AudioFilesIdJSONRequestBody_Map(t *testing.T) {
	tests := []struct {
		name string
		pa   PatchV1AudioFilesIdJSONRequestBody
		want map[string]string
	}{
		{
			name: "all fields filled",
			pa: PatchV1AudioFilesIdJSONRequestBody{
				Album:        "Moving Pictures",
				Genre:        "Rock",
				Performer:    "Rush",
				RecordedDate: helpers.ToPtr(1981),
				Title:        "Tom Sawyer",
				TrackNumber:  helpers.ToPtr(1),
			},
			want: map[string]string{
				"album":  "Moving Pictures",
				"artist": "Rush",
				"date":   "1981",
				"genre":  "Rock",
				"title":  "Tom Sawyer",
				"track":  "1",
			},
		},
		{
			name: "only genre",
			pa: PatchV1AudioFilesIdJSONRequestBody{
				Genre: "Rock",
			},
			want: map[string]string{
				"genre": "Rock",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.pa.Map()
			assert.Equal(t, tt.want, got)
		})
	}
}
