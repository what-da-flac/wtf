package interfaces

import "github.com/what-da-flac/wtf/openapi/gen/golang"

//go:generate moq -out ../../mocks/timer.go -pkg mocks . Repository
type Repository interface {
	InsertFile(file *golang.File) error
}
