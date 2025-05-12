package ifaces

import (
	"github.com/what-da-flac/wtf/openapi/gen/golang"
)

//go:generate moq -out ../../mocks/repository.go -pkg mocks . Repository
type Repository interface {
	FindByHash(hash string) (*golang.AudioFile, error)
	InsertAudioFile(file *golang.AudioFile) error
	SelectAudioFile(id string) (*golang.AudioFile, error)
	UpdateAudioFile(id string, values map[string]any) error
}
