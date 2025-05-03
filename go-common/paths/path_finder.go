package paths

import (
	"io"
	"os"
	"path/filepath"

	"github.com/what-da-flac/wtf/go-common/ifaces"
)

type PathFinder struct {
	baseDir    string
	identifier ifaces.Identifier
}

func NewPathFinder(identifier ifaces.Identifier, baseDir string) *PathFinder {
	return &PathFinder{
		identifier: identifier,
		baseDir:    baseDir,
	}
}

func (x *PathFinder) Path() string {
	return x.baseDir
}

func (x *PathFinder) SaveSteam(r io.Reader) (string, error) {
	if err := os.MkdirAll(x.baseDir, os.ModePerm); err != nil {
		return "", err
	}
	filename := x.identifier.UUIDv4()
	filename = filepath.Join(x.baseDir, filename)
	dst, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer func() { _ = dst.Close() }()
	_, err = io.Copy(dst, r)
	return filename, err
}
