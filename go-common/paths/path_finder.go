package paths

import (
	"io"
	"os"
	"path/filepath"

	"github.com/what-da-flac/wtf/openapi/gen/golang"
)

type PathFinder struct {
	baseDir  string
	pathName golang.PathName
}

func NewPathFinder(baseDir string, pathName golang.PathName) *PathFinder {
	return &PathFinder{
		baseDir:  baseDir,
		pathName: pathName,
	}
}

func (x *PathFinder) Path() golang.PathName {
	return x.pathName
}

func (x *PathFinder) Save(r io.Reader, key string) error {
	if err := os.MkdirAll(x.baseDir, os.ModePerm); err != nil {
		return err
	}
	filename := filepath.Join(x.baseDir, key)
	dst, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func() { _ = dst.Close() }()
	_, err = io.Copy(dst, r)
	return err
}
