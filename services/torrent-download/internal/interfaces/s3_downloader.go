package interfaces

import (
	"io"
)

//go:generate moq -out ../../mocks/s3_downloader.go -pkg mocks . S3Downloader
type S3Downloader interface {
	Download(w io.WriterAt, bucket, key string) error
}
