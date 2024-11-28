package processors

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/what-da-flac/wtf/go-common/env"
	"github.com/what-da-flac/wtf/go-common/ifaces"
	"github.com/what-da-flac/wtf/openapi/models"
	"github.com/what-da-flac/wtf/services/torrent-download/internal/interfaces"
	"github.com/what-da-flac/wtf/services/torrent-download/internal/model"
)

type Processor struct {
	logger            ifaces.Logger
	s3Downloader      interfaces.S3Downloader
	torrentDownloader interfaces.TorrentDownloader
}

func NewProcessor(logger ifaces.Logger, torrentDownloader interfaces.TorrentDownloader,
	s3Downloader interfaces.S3Downloader) *Processor {
	return &Processor{
		logger:            logger,
		s3Downloader:      s3Downloader,
		torrentDownloader: torrentDownloader,
	}
}

func (x *Processor) Process(
	torrent *models.Torrent,
	config *env.Config, workingDir string) (*time.Duration, error) {
	// validate incoming torrent
	if err := x.validateTorrent(torrent); err != nil {
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
	// create torrent destination file
	torrentFilename := filepath.Join(torrentDir, torrent.Filename)
	torrentFile, err := os.Create(torrentFilename)
	if err != nil {
		return nil, err
	}
	// download torrent file from s3
	if err = x.downloadTorrentFromS3(torrentFile, torrent); err != nil {
		return nil, err
	}
	targetDir := filepath.Join(config.Volumes.Downloads.String(), torrent.Id)
	if err := os.MkdirAll(targetDir, os.ModePerm); err != nil {
		return nil, err
	}

	// download torrent contents
	if err = x.downloadTorrentContents(config.Downloads.Timeout, targetDir, torrentFilename); err != nil {
		return nil, err
	}
	elapsed := time.Since(start)
	return &elapsed, nil
}

func (x *Processor) downloadTorrentFromS3(w io.WriteCloser, torrent *models.Torrent) error {
	x.logger.Infof("starting downloading torrent from s3: %s", torrent.Filename)
	defer func() { _ = w.Close() }()
	if err := x.s3Downloader.Download(w, env.BucketTorrentParsed.String(), torrent.Filename); err != nil {
		x.logger.Errorf("failed to download torrent from s3: %s", torrent.Filename)
		return err
	}
	return nil
}

func (x *Processor) downloadTorrentContents(
	timeout time.Duration,
	targetDir, torrentFilename string) error {
	var lastProgress float64
	interval := time.Second * 5
	if err := x.torrentDownloader.Start(); err != nil {
		return err
	}
	// clean up resources for next execution
	defer func() {
		_ = x.torrentDownloader.Stop()
	}()
	if err := x.torrentDownloader.AddTorrent(targetDir, torrentFilename,
		func(line *model.TorrentLine) {
			if line.Percent != lastProgress {
				// TODO: publish to torrent-info
				lastProgress = line.Percent
				x.logger.Info(line.String())
			}
		},
	); err != nil {
		return err
	}
	if !x.torrentDownloader.WaitForDownload(timeout, interval) {
		// if download was not successful, remove all files and torrents
		if err := x.torrentDownloader.RemoveAll(); err != nil {
			return err
		}
		return fmt.Errorf("torrent download timed out after: %v", timeout)
	}
	if err := x.torrentDownloader.ClearTorrents(); err != nil {
		return err
	}
	return nil
}

func (x *Processor) validateTorrent(torrent *models.Torrent) error {
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
