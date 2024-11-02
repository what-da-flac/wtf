package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/what-da-flac/wtf/go-common/ifaces"
	"github.com/what-da-flac/wtf/lambdas/torrent-download/internal/downloaders"
	"go.uber.org/zap"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/what-da-flac/wtf/go-common/amazon"
	"github.com/what-da-flac/wtf/lambdas/torrent-download/internal/environment"
	"github.com/what-da-flac/wtf/openapi/models"
)

var (
	logger ifaces.Logger
)

func main() {
	lambda.Start(handler)
}

// handler converts magnet link to torrent file, extracts info and files,
// and sends sqs message with the information.
func handler(_ context.Context, sqsEvent *events.SQSEvent) error {
	zl, err := zap.NewDevelopment()
	if err != nil {
		return err
	}
	logger = zl.Sugar()
	config := environment.New()
	// loop over messages received
	for _, record := range sqsEvent.Records {
		payload := &models.Torrent{}
		if err := json.Unmarshal([]byte(record.Body), payload); err != nil {
			log.Println(err)
			return err
		}
		if err := process(config, payload); err != nil {
			log.Println(err)
			// TODO: if above fails, send a message to another queue to deal with failed torrents
			// ignoring it for the time being
			return nil
		}
	}
	return nil
}

func process(config *environment.Config, torrent *models.Torrent) error {
	// validate incoming torrent
	if err := validateTorrent(torrent); err != nil {
		return err
	}
	// create aws session object
	awsSession := amazon.NewAWSSessionFromEnvironment()
	if err := awsSession.Build(); err != nil {
		return err
	}
	sess := awsSession.Session()
	// base dir must be /tmp since lambdas cannot write anywhere else
	tmpDir := os.TempDir()
	// download torrent from s3
	torrentDir := filepath.Join(tmpDir, "torrents")
	if err := os.MkdirAll(torrentDir, 0700); err != nil {
		return err
	}
	// clean up resources for next lambda execution
	defer func() { _ = os.RemoveAll(torrentDir) }()
	torrentFilename, err := downloadTorrentFromS3(config, sess, torrent, torrentDir)
	if err != nil {
		return err
	}
	targetDir := filepath.Join(tmpDir, "downloads")
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return err
	}
	// clean up resources for next lambda execution
	defer func() { _ = os.RemoveAll(targetDir) }()

	// download torrent contents
	if err = downloadTorrent(logger, config, targetDir, *torrentFilename); err != nil {
		return err
	}
	// upload all resulting files to s3
	return uploadResultToS3(sess, config, torrent, targetDir)
}

func uploadResultToS3(
	sess *session.Session, config *environment.Config,
	torrent *models.Torrent, targetDir string) error {
	bucket := config.BucketDownloads
	uploadFn := func(key, filename string) error {
		info, err := os.Stat(filename)
		if err != nil {
			return err
		}
		file, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer func() { _ = file.Close() }()
		return amazon.Upload(sess, file, bucket, key, amazon.Content{
			// TODO: set all fields
			// ContentDisposition: "",
			// ContentEncoding:    "",
			ContentLanguage: "en",
			ContentLength:   info.Size(),
			// ContentType:        "",
		})
	}
	// loop over all files in targetDir and send to s3
	return filepath.Walk(targetDir, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if err != nil {
			// ignoring files that cannot be processed
			return nil
		}
		basePath := strings.TrimPrefix(path, targetDir)
		key := filepath.Join(torrent.Id, basePath)
		return uploadFn(key, path)
	})
}

func downloadTorrentFromS3(config *environment.Config, sess *session.Session,
	torrent *models.Torrent, targetDir string) (*string, error) {
	// torrent file from s3 is downloaded to /tmp/current.torrent
	const targetTorrentFilename = "current.torrent"
	filename := filepath.Join(targetDir, targetTorrentFilename)
	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	defer func() { _ = file.Close() }()
	if err := amazon.Download(sess, file, config.BucketDownloads, torrent.Filename); err != nil {
		return nil, err
	}
	return &filename, nil
}

func downloadTorrent(logger ifaces.Logger, config *environment.Config,
	targetDir, torrentFilename string) error {
	interval := time.Second * 5
	downloader := downloaders.NewTorrentDownloader(logger)
	if err := downloader.Start(); err != nil {
		return err
	}
	if err := downloader.AddTorrent(targetDir, torrentFilename); err != nil {
		return err
	}
	if downloader.WaitForDownload(config.Timeout, interval) {
		return nil
	}
	return fmt.Errorf("could not complete torrent download before timeout")
}

func validateTorrent(torrent *models.Torrent) error {
	if torrent == nil {
		return fmt.Errorf("torrent is nil")
	}
	if torrent.Id == "" {
		return fmt.Errorf("torrent id is empty")
	}
	if torrent.Filename == "" {
		return fmt.Errorf("torrent filename is empty")
	}
	return nil
}
