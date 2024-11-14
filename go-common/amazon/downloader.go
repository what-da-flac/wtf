package amazon

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Downloader struct {
	downloader *s3manager.Downloader
}

func NewDownloader(sess *session.Session) *Downloader {
	return &Downloader{
		downloader: s3manager.NewDownloader(sess),
	}
}

func (x *Downloader) Download(w io.WriterAt, bucket, key string) error {
	input := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}
	_, err := x.downloader.Download(w, input)
	return err
}
