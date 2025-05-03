package interfaces

import (
	"io"

	"github.com/what-da-flac/wtf/openapi/gen/golang"
)

//go:generate moq -out ../../mocks/file_storage.go -pkg mocks . FileStorage
type FileStorage interface {
	// Save stores file content into persistent storage, and returns path to filename.
	Save(f *golang.File, file io.Reader) (string, error)
}
