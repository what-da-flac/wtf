package interfaces

import (
	"io"

	"github.com/what-da-flac/wtf/openapi/gen/golang"
)

//go:generate moq -out ../../mocks/path_finder.go -pkg mocks . PathFinder
type PathFinder interface {
	// Path returns absolute path.
	Path() golang.PathName

	// Save writes reader to file using key as its name.
	Save(r io.Reader, key string) error
}
