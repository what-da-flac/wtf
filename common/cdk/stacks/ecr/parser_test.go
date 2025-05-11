package ecr

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
			name: "sample-ecr",
			args: args{
				filename: filepath.Join("test-data", "sample-ecr.yaml"),
			},
			want: []*Model{
				{
					EmptyOnDelete:   true,
					Mutable:         true,
					Name:            "email-service",
					RemoveOnDestroy: true,
				},
				{
					Name:            "file-manager-service",
					EmptyOnDelete:   true,
					Mutable:         true,
					RemoveOnDestroy: true,
					UseDefaults:     true,
				},
				{
					Name:            "identity-manager-service",
					EmptyOnDelete:   true,
					Mutable:         true,
					RemoveOnDestroy: true,
					UseDefaults:     true,
				},
				{
					Name:            "user-service",
					EmptyOnDelete:   true,
					Mutable:         true,
					RemoveOnDestroy: true,
					UseDefaults:     true,
				},
				{
					Name:            "webapp",
					EmptyOnDelete:   true,
					Mutable:         true,
					RemoveOnDestroy: true,
					UseDefaults:     true,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.filename)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			if err != nil {
				return
			}
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}
