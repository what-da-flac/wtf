package files

import (
	"embed"
	"io/fs"
	"path/filepath"
)

// FSResult is the outcome of processing files within a FSPath struct.
type FSResult struct {
	Path  string
	Entry fs.DirEntry
	File  *embed.FS
}

// FullPath returns the full readable path to specific embedded file.
func (x *FSResult) FullPath() string {
	return filepath.Join(x.Path, x.Entry.Name())
}
