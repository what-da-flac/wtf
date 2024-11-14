package processors

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/what-da-flac/wtf/go-common/amazon"
	"github.com/what-da-flac/wtf/go-common/env"
	"github.com/what-da-flac/wtf/go-common/ifaces"
	"github.com/what-da-flac/wtf/openapi/models"
	"github.com/what-da-flac/wtf/services/torrent-download/internal/downloaders"
)

func Process(sess *session.Session, logger ifaces.Logger,
	torrent *models.Torrent, config *env.Config) (*time.Duration, error) {
	// validate incoming torrent
	if err := validateTorrent(torrent); err != nil {
		return nil, err
	}
	start := time.Now()
	// setting working directory
	tmpDir := os.TempDir()
	// download torrent from s3
	torrentDir := filepath.Join(tmpDir, "torrents")
	if err := os.MkdirAll(torrentDir, 0700); err != nil {
		return nil, err
	}
	// clean up resources for next  execution
	defer func() { _ = os.RemoveAll(torrentDir) }()
	torrentFilename, err := downloadTorrentFromS3(logger, sess, torrent, torrentDir)
	if err != nil {
		return nil, err
	}
	targetDir := filepath.Join(config.Volumes.Downloads.String(), torrent.Id)
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return nil, err
	}
	// clean up resources for next lambda execution
	defer func() { _ = os.RemoveAll(targetDir) }()

	// download torrent contents
	if err = downloadTorrentContents(logger, config.Downloads.Timeout, targetDir, *torrentFilename); err != nil {
		return nil, err
	}
	elapsed := time.Since(start)
	return &elapsed, nil
}

func downloadTorrentFromS3(logger ifaces.Logger, sess *session.Session,
	torrent *models.Torrent, targetDir string) (*string, error) {
	// torrent file from s3 is downloaded to /tmp/current.torrent
	filename := filepath.Join(targetDir, torrent.Filename)
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
	defer func() {
		_ = downloader.Stop()
	}()
	if err := downloader.AddTorrent(targetDir, torrentFilename); err != nil {
		return err
	}
	if !downloader.WaitForDownload(timeout, interval) {
		// if download was not successful, remove all files and torrents
		return downloader.RemoveAll()
	}
	if err := downloader.ClearTorrents(); err != nil {
		return err
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
