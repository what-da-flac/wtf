package ifaces

import (
	"github.com/what-da-flac/wtf/openapi/gen/golang"
)

//go:generate moq -out ../../mocks/repository.go -pkg mocks . Repository
type Repository interface {
	InsertAudioFile(file *golang.AudioFile) error
}
