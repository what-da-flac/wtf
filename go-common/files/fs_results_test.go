package files

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFSResults_WriteDir(t *testing.T) {
	const format = "20060102150405.000000000"
	now := time.Now()
	targetDir := filepath.Join(os.TempDir(), now.Format(format))
	t.Logf("target directory: %s", targetDir)
	results, err := Join(NewFSPath(filesFs1), NewFSPath(filesFs2))
	assert.NoError(t, err)
	err = results.WriteDir(targetDir)
	assert.NoError(t, err)
	err = os.RemoveAll(targetDir)
	assert.NoError(t, err)
}

func TestFSResults_WriteUniqueFiles(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name    string
		x       FSResults
		args    args
		wantErr error
	}{
		{
			name: "happy path",
			x: func() FSResults {
				results, err := Join(NewFSPath(filesFs1), NewFSPath(filesFs2))
				assert.NoError(t, err)
				return results
			}(),
			args: args{
				dir: func() string {
					const format = "20060102150405.000000000"
					now := time.Now()
					dir := filepath.Join(os.TempDir(), now.Format(format))
					t.Log("output dir: ", dir)
					return dir
				}(),
			},
			wantErr: nil,
		},
		{
			name: "dups error",
			x: func() FSResults {
				results, err := Join(NewFSPath(filesFs1), NewFSPath(filesFs2), NewFSPath(filesFs2))
				assert.NoError(t, err)
				return results
			}(),
			args: args{
				dir: func() string {
					const format = "20060102150405.000000000"
					now := time.Now()
					dir := filepath.Join(os.TempDir(), now.Format(format))
					return dir
				}(),
			},
			wantErr: fmt.Errorf("file already exists: 20240101000000_user-tables.down.sql"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.x.WriteUniqueFiles(tt.args.dir)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
