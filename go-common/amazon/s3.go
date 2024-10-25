package amazon

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
)

type S3 struct {
	bucket string
	sess   *session.Session
}

func NewS3(sess *session.Session, bucket string) *S3 {
	return &S3{
		bucket: bucket,
		sess:   sess,
	}
}

func (x *S3) Bucket() string { return x.bucket }

func (x *S3) Download(file *os.File, key string) error {
	return Download(x.sess, file, x.Bucket(), key)
}

func (x *S3) DownloadWithBucket(file *os.File, bucket, key string) error {
	return Download(x.sess, file, bucket, key)
}

func (x *S3) Upload(file *os.File, key string, content Content) error {
	return Upload(x.sess, file, x.Bucket(), key, content)
}

func (x *S3) UploadWithBucket(file *os.File, bucket, key string, content Content) error {
	return Upload(x.sess, file, bucket, key, content)
}

func (x *S3) PreSignDownload(key string, expiration time.Duration) (*string, error) {
	return PreSignDownload(x.sess, x.Bucket(), key, expiration)
}

func (x *S3) Info(bucket, key string) (*S3FileInfo, error) {
	return Info(x.sess, bucket, key)
}

func (x *S3) DownloadPreSignedURL(w io.Writer, presignedUrl string) error {
	// intentionally no timeout is being set
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, presignedUrl, nil)
	if err != nil {
		return err
	}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func() { _ = res.Body.Close() }()
	if expected, actual := http.StatusOK, res.StatusCode; expected != actual {
		data, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("expected status code: %d, got: %d, message: %s", expected, actual, string(data))
	}
	_, err = io.Copy(w, res.Body)
	return err
}

func (x *S3) Copy(srcBucket, srcKey, dstBucket, dstKey string) error {
	return Copy(x.sess, srcBucket, srcKey, dstBucket, dstKey)
}

func (x *S3) Remove(bucket, key string) error {
	return Remove(x.sess, bucket, key)
}
