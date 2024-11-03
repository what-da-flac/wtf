package uploaders

import (
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/what-da-flac/wtf/go-common/amazon"
)

type Uploader struct {
	sess *session.Session
}

func NewUploader(sess *session.Session) *Uploader {
	return &Uploader{
		sess: sess,
	}
}

func (x *Uploader) Upload(file *os.File, bucket, key string, content amazon.Content) error {
	return amazon.Upload(x.sess, file, bucket, key, content)
}
