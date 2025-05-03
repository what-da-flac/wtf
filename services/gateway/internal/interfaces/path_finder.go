package interfaces

import (
	"io"
)

//go:generate moq -out ../../mocks/path_finder.go -pkg mocks . PathFinder
type PathFinder interface {
	// Path returns absolute path.
	Path() string

	// SaveSteam writes reader to filename. Returns resulting path to saved file.
	SaveSteam(r io.Reader) (string, error)
}
