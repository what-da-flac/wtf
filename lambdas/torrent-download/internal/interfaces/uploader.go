package interfaces

import (
	"os"

	"github.com/what-da-flac/wtf/go-common/amazon"
)

//go:generate moq -out ../../mocks/uploader.go -pkg mocks . Uploader
type Uploader interface {
	Upload(file *os.File, bucket, key string, content amazon.Content) error
}
