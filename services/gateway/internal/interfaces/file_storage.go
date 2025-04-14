package interfaces

import (
	"io"

	"github.com/what-da-flac/wtf/openapi/domains"
)

//go:generate moq -out ../../mocks/file_storage.go -pkg mocks . FileStorage
type FileStorage interface {
	// Save stores file content into persistent storage
	Save(f *domains.File, file io.Reader) error
}
