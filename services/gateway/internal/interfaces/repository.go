package interfaces

import (
	"github.com/what-da-flac/wtf/openapi/domains"
)

//go:generate moq -out ../../mocks/repository.go -pkg mocks . Repository
type Repository interface {
	InsertFile(file *domains.File) error
}
