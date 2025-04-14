package storages

import (
	"io"
	"os"
	"path/filepath"

	"github.com/what-da-flac/wtf/openapi/domains"
)

type Local struct {
	dirPath string
}

func NewLocal(dirPath string) *Local {
	return &Local{
		dirPath: dirPath,
	}
}

func (x *Local) Save(f *domains.File, file io.Reader) error {
	filename := x.Filename(*f)
	if err := os.MkdirAll(x.dirPath, os.ModePerm); err != nil {
		return err
	}
	dst, err := os.Create(filepath.Join(x.dirPath, filename))
	if err != nil {
		return err
	}
	defer func() { _ = dst.Close() }()
	_, err = io.Copy(dst, file)
	return err
}

func (x *Local) Filename(f domains.File) string {
	base := filepath.Base(f.Filename)
	ext := filepath.Ext(base)
	return f.Id + ext
}
