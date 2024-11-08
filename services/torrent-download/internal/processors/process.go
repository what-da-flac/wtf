package processors

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/what-da-flac/wtf/go-common/amazon"
	"github.com/what-da-flac/wtf/go-common/env"
	"github.com/what-da-flac/wtf/go-common/ifaces"
	"github.com/what-da-flac/wtf/openapi/models"
	"github.com/what-da-flac/wtf/services/torrent-download/internal/downloaders"
)

func Process(sess *session.Session, logger ifaces.Logger, downloadTimeout time.Duration, torrent *models.Torrent) error {
	// validate incoming torrent
	if err := validateTorrent(torrent); err != nil {
		return err
	}
	// base dir must be /tmp since lambdas cannot write anywhere else
	tmpDir := os.TempDir()
	// download torrent from s3
	torrentDir := filepath.Join(tmpDir, "torrents")
	if err := os.MkdirAll(torrentDir, 0700); err != nil {
		return err
	}
	// clean up resources for next lambda execution
	defer func() { _ = os.RemoveAll(torrentDir) }()
	torrentFilename, err := downloadTorrentFromS3(logger, sess, torrent, torrentDir)
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
	if err = downloadTorrentContents(logger, downloadTimeout, targetDir, *torrentFilename); err != nil {
		return err
	}
	// upload all resulting files to s3
	return uploadResultToS3(sess, logger, torrent, targetDir)
}

func uploadResultToS3(
	sess *session.Session,
	logger ifaces.Logger,
	torrent *models.Torrent, targetDir string) error {
	const slash = "/"
	bucket := env.BucketTorrentParsed.String()
	// targetDir requires a trailing slash
	uploadFn := func(key, filename string) error {
		logger.Infof("starting uploading file: %s", key)
		info, err := os.Stat(filename)
		if err != nil {
			logger.Errorf("failed to stat file: %s", filename)
			return err
		}
		file, err := os.Open(filename)
		if err != nil {
			logger.Errorf("failed to open file: %s", filename)
			return err
		}
		defer func() { _ = file.Close() }()
		if err = amazon.Upload(sess, file, bucket, key, amazon.Content{
			//ContentDisposition: "",
			//ContentEncoding:    "",
			ContentLanguage: "en",
			ContentLength:   info.Size(),
			//ContentType:        "",
		}); err != nil {
			logger.Errorf("failed to upload file: %s", filename)
		}
		return nil
	}
	// TODO: instead of walking directory, loop over torrent select files
	// loop over all files in targetDir and send to s3
	return filepath.Walk(targetDir, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if err != nil {
			// ignoring files that cannot be processed
			return nil
		}
		// remove targetDir from path
		basePath := strings.TrimPrefix(path, targetDir+slash)
		// remove torrent internal directory from path
		basePath = strings.TrimPrefix(basePath, torrent.Name+slash)
		// use only the id as root for torrent contents
		key := filepath.Join(torrent.Id, basePath)
		return uploadFn(key, path)
	})
}

func downloadTorrentFromS3(logger ifaces.Logger, sess *session.Session,
	torrent *models.Torrent, targetDir string) (*string, error) {
	// torrent file from s3 is downloaded to /tmp/current.torrent
	const targetTorrentFilename = "current.torrent"
	filename := filepath.Join(targetDir, targetTorrentFilename)
	file, err := os.Create(filename)
	if err != nil {
		logger.Errorf("failed to create file: %s", filename)
		return nil, err
	}
	defer func() { _ = file.Close() }()
	logger.Infof("starting downloading torrent from s3: %s", torrent.Filename)
	if err := amazon.Download(sess, file, env.BucketTorrentParsed.String(), torrent.Filename); err != nil {
		logger.Errorf("failed to download torrent from s3: %s", torrent.Filename)
		return nil, err
	}
	return &filename, nil
}

func downloadTorrentContents(logger ifaces.Logger, timeout time.Duration,
	targetDir, torrentFilename string) error {
	interval := time.Second * 5
	daemonTimeout := time.Second * 10
	downloader := downloaders.NewTorrentDownloader(logger, daemonTimeout)
	if err := downloader.Start(); err != nil {
		return err
	}
	// clean up resources for next execution
	defer func() { _ = downloader.ClearAll() }()
	if err := downloader.AddTorrent(targetDir, torrentFilename); err != nil {
		return err
	}
	if downloader.WaitForDownload(timeout, interval) {
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