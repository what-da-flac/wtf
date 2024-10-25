package amazon

import (
	"os"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Content struct {
	ContentDisposition string
	ContentEncoding    string
	ContentLanguage    string
	ContentLength      int64
	ContentType        string
}

func (x Content) putObjectInput() *s3.PutObjectInput {
	return &s3.PutObjectInput{
		ContentDisposition: aws.String(x.ContentDisposition),
		ContentEncoding:    aws.String(x.ContentEncoding),
		ContentLanguage:    aws.String(x.ContentLanguage),
		ContentLength:      aws.Int64(x.ContentLength),
		ContentType:        aws.String(x.ContentType),
	}
}

func Download(sess *session.Session, file *os.File, bucket, key string) error {
	downloader := s3manager.NewDownloader(sess)
	input := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}
	if _, err := downloader.Download(file, input); err != nil {
		return err
	}
	return nil
}

func Upload(sess *session.Session, file *os.File, bucket, key string, content Content) error {
	svc := s3.New(sess)
	input := content.putObjectInput()
	input.Bucket = aws.String(bucket)
	input.Key = aws.String(key)
	input.Body = file
	_, err := svc.PutObject(input)
	return err
}

func Copy(sess *session.Session, srcBucket, srcKey, dstBucket, dstKey string) error {
	svc := s3.New(sess)
	input := &s3.CopyObjectInput{
		Bucket:     aws.String(dstBucket),
		CopySource: aws.String(filepath.Join(srcBucket, srcKey)),
		Key:        aws.String(dstKey),
	}
	_, err := svc.CopyObject(input)
	return err
}

func Remove(sess *session.Session, bucket, key string) error {
	svc := s3.New(sess)
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}
	_, err := svc.DeleteObject(input)
	return err
}

func PreSignDownload(sess *session.Session, bucket, key string, expiration time.Duration) (*string, error) {
	svc := s3.New(sess)
	input := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}
	req, _ := svc.GetObjectRequest(input)
	uri, err := req.Presign(expiration)
	if err != nil {
		return nil, err
	}
	return &uri, nil
}

func Info(sess *session.Session, bucket, key string) (*S3FileInfo, error) {
	svc := s3.New(sess)
	input := &s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}
	out, err := svc.HeadObject(input)
	if err != nil {
		return nil, err
	}
	return &S3FileInfo{
		Bucket:        bucket,
		ContentLength: *out.ContentLength,
		ContentType:   *out.ContentType,
		Key:           key,
		LastModified:  *out.LastModified,
	}, nil
}
