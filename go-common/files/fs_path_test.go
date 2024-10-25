package files

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJoin(t *testing.T) {
	joined, err := Join(NewFSPath(filesFs1), NewFSPath(filesFs2))
	assert.NoError(t, err)
	filtered := joined.Filter(func(r *FSResult) bool {
		return !r.Entry.IsDir() && !strings.HasPrefix(r.Entry.Name(), ".")
	})
	sorted := filtered.Sort()
	expected := []struct {
		name   string
		length int
	}{
		{
			name:   "test-data/2/20240101000000_user-tables.down.sql",
			length: 0,
		},
		{
			name:   "test-data/2/20240101000000_user-tables.up.sql",
			length: 153,
		},
		{
			name:   "test-data/1/20240325012625_email-tables.down.sql",
			length: 0,
		},
		{
			name:   "test-data/1/20240325012625_email-tables.up.sql",
			length: 952,
		},
	}
	for i, r := range sorted {
		data, err := r.File.ReadFile(r.FullPath())
		assert.NoError(t, err)
		expect := expected[i]
		assert.Equal(t, r.FullPath(), expect.name)
		assert.Equal(t, len(data), expect.length)
	}
}

func TestPathsFS_Results(t *testing.T) {
	var pathFS = FSPath{
		filesFs,
	}
	results, err := pathFS.Entries()
	assert.NoError(t, err)
	filtered := results.Filter(func(r *FSResult) bool {
		// ignore directories
		if r.Entry.IsDir() {
			return false
		}
		// ignore hidden files
		if strings.HasPrefix(r.Entry.Name(), ".") {
			return false
		}
		return true
	})
	assert.Len(t, filtered, 4)
	sorted := filtered.Sort()
	for _, r := range sorted {
		t.Logf("path: %s name: %s", r.Path, r.Entry.Name())
	}
}
