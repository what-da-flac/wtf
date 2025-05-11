package s3

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    []*Model
		wantErr bool
	}{
		{
			name: "happy-path.yaml",
			args: args{
				filename: "happy-path.yaml",
			},
			want: []*Model{
				{
					Name:            "wtf.torrent-parsed",
					RemoveOnDestroy: true,
				},
				{
					Name:            "wtf.torrent-downloads",
					RemoveOnDestroy: true,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filename := filepath.Join("test-data", tt.args.filename)
			got, err := Parse(filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("fromYAML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
