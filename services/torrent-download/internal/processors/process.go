package processors

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/what-da-flac/wtf/go-common/env"
	"github.com/what-da-flac/wtf/go-common/ifaces"
	"github.com/what-da-flac/wtf/openapi/models"
	"github.com/what-da-flac/wtf/services/torrent-download/internal/interfaces"
)

func Process(
	logger ifaces.Logger, torrentDownloader interfaces.TorrentDownloader,
	s3Downloader interfaces.S3Downloader, torrent *models.Torrent,
	config *env.Config, workingDir string) (*time.Duration, error) {
	// validate incoming torrent
	if err := validateTorrent(torrent); err != nil {
		return nil, err
	}
	start := time.Now()
	// download torrent from s3
	torrentDir := filepath.Join(workingDir, "torrents")
	if err := os.MkdirAll(torrentDir, 0700); err != nil {
		return nil, err
	}
	// clean up resources for next  execution
	defer func() { _ = os.RemoveAll(torrentDir) }()
	torrentFilename, err := downloadTorrentFromS3(logger, s3Downloader, torrent, torrentDir)
	if err != nil {
		return nil, err
	}
	targetDir := filepath.Join(config.Volumes.Downloads.String(), torrent.Id)
	if err := os.MkdirAll(targetDir, os.ModePerm); err != nil {
		return nil, err
	}

	// download torrent contents
	if err = downloadTorrentContents(torrentDownloader, config.Downloads.Timeout, targetDir, *torrentFilename); err != nil {
		return nil, err
	}
	elapsed := time.Since(start)
	return &elapsed, nil
}

func downloadTorrentFromS3(logger ifaces.Logger, downloader interfaces.S3Downloader,
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
	if err := downloader.Download(file, env.BucketTorrentParsed.String(), torrent.Filename); err != nil {
		logger.Errorf("failed to download torrent from s3: %s", torrent.Filename)
		return nil, err
	}
	return &filename, nil
}

func downloadTorrentContents(downloader interfaces.TorrentDownloader, timeout time.Duration,
	targetDir, torrentFilename string) error {
	interval := time.Second * 5
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
		if err := downloader.RemoveTorrentsAndFiles(); err != nil {
			return err
		}
		return fmt.Errorf("torrent download timed out after: %v", timeout)
	}
	if err := downloader.RemoveTorrentsLeaveFiles(); err != nil {
		return err
	}
	return nil
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
