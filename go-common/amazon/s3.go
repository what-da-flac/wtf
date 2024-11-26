package amazon

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"golang.org/x/net/context"
)

type Content struct {
	ContentDisposition string
	ContentEncoding    string
	ContentLanguage    string
	ContentLength      int64
	ContentType        string
}

type S3 struct {
	client *s3.Client
}

func NewS3(cfg *aws.Config) *S3 {
	return &S3{
		client: s3.NewFromConfig(*cfg),
	}
}

func (x *S3) Download(w io.Writer, bucket, key string) error {
	input := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}
	res, err := x.client.GetObject(context.TODO(), input)
	if err != nil {
		return err
	}
	defer func() { _ = res.Body.Close() }()
	_, err = io.Copy(w, res.Body)
	return err
}

func (x *S3) Upload(r io.Reader, bucket, key string, content Content) error {
	length := content.ContentLength
	if length == 0 {
		return fmt.Errorf("missing contentLength")
	}
	input := &s3.PutObjectInput{
		Bucket:        aws.String(bucket),
		Key:           aws.String(key),
		Body:          r,
		ContentLength: aws.Int64(length),
	}
	if content.ContentType != "" {
		input.ContentType = aws.String(content.ContentType)
	}
	if content.ContentDisposition != "" {
		input.ContentDisposition = aws.String(content.ContentDisposition)
	}
	if content.ContentEncoding != "" {
		input.ContentEncoding = aws.String(content.ContentEncoding)
	}
	if content.ContentLanguage != "" {
		input.ContentLanguage = aws.String(content.ContentLanguage)
	}
	_, err := x.client.PutObject(context.TODO(), input)
	return err
}

func (x *S3) PreSignDownload(bucket, key string, expiration time.Duration) (*string, error) {
	client := s3.NewPresignClient(x.client)
	url, err := client.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}, func(o *s3.PresignOptions) {
		o.Expires = expiration
	})
	if err != nil {
		return nil, err
	}
	return &url.URL, nil
}

func (x *S3) Info(bucket, key string) (*s3.HeadObjectOutput, error) {
	input := &s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}
	return x.client.HeadObject(context.TODO(), input)
}

func (x *S3) DownloadPreSignedURL(w io.Writer, url string) error {
	// intentionally no timeout is being set
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
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
	srcBucket = strings.TrimPrefix(srcBucket, "/")
	input := &s3.CopyObjectInput{
		Bucket:     aws.String(dstBucket),
		CopySource: aws.String(srcBucket + "/" + srcKey),
		Key:        aws.String(dstKey),
	}
	_, err := x.client.CopyObject(context.TODO(), input)
	return err
}

func (x *S3) Remove(bucket, key string) error {
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}
	_, err := x.client.DeleteObject(context.TODO(), input)
	return err
}
