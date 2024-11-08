package pipes

import (
	"fmt"
	"io"
	"os"
)

var ErrNoPipe = fmt.Errorf("no pipe")

// ReadStdin tries to read from stdin pipe, if empty returns an error.
// If not empty, returns the reader that contains stdin data.
func ReadStdin() (io.Reader, error) {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return nil, err
	}
	if fi.Mode()&os.ModeNamedPipe == 0 {
		return nil, ErrNoPipe
	}
	return os.Stdin, nil
}
