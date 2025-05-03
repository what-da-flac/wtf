package storages

import (
	"io"
	"os"
	"path/filepath"

	"github.com/what-da-flac/wtf/openapi/gen/golang"
)

type Local struct {
	dirPath string
}

func NewLocal(dirPath string) *Local {
	return &Local{
		dirPath: dirPath,
	}
}

func (x *Local) Save(f *golang.File, file io.Reader) (string, error) {
	filename := x.Filename(f)
	if err := os.MkdirAll(x.dirPath, os.ModePerm); err != nil {
		return "", err
	}
	newFilename := filepath.Join(x.dirPath, filename)
	dst, err := os.Create(newFilename)
	if err != nil {
		return "", err
	}
	defer func() { _ = dst.Close() }()
	_, err = io.Copy(dst, file)
	return newFilename, err
}

func (x *Local) Filename(f *golang.File) string {
	base := filepath.Base(f.Filename)
	ext := filepath.Ext(base)
	return f.Id + ext
}
