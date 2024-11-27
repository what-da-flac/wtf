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
	"github.com/what-da-flac/wtf/services/torrent-download/internal/model"
)

type Processor struct {
	config            *env.Config
	logger            ifaces.Logger
	s3Downloader      interfaces.S3Downloader
	torrentDownloader interfaces.TorrentDownloader
	workingDir        string
}

func NewProcessor(
	config *env.Config,
	logger ifaces.Logger,
	s3Downloader interfaces.S3Downloader,
	torrentDownloader interfaces.TorrentDownloader,
	workingDir string) *Processor {
	return &Processor{
		config:            config,
		logger:            logger,
		s3Downloader:      s3Downloader,
		torrentDownloader: torrentDownloader,
		workingDir:        workingDir,
	}
}

func (x *Processor) Process(torrent *models.Torrent) (*time.Duration, error) {
	// validate incoming torrent
	if err := x.validateTorrent(torrent); err != nil {
		return nil, err
	}
	start := time.Now()
	// download torrent from s3
	torrentDir := filepath.Join(x.workingDir, "torrents")
	if err := os.MkdirAll(torrentDir, 0700); err != nil {
		return nil, err
	}
	// clean up resources for next  execution
	defer func() { _ = os.RemoveAll(torrentDir) }()
	if _, err := x.downloadTorrentFromS3(torrent); err != nil {
		return nil, err
	}
	targetDir := filepath.Join(x.config.Volumes.Downloads.String(), torrent.Id)
	if err := os.MkdirAll(targetDir, os.ModePerm); err != nil {
		return nil, err
	}

	// download torrent contents
	if err := x.downloadTorrentContents(torrent, x.config.Downloads.Timeout); err != nil {
		return nil, err
	}
	elapsed := time.Since(start)
	return &elapsed, nil
}

func (x *Processor) downloadTorrentFromS3(torrent *models.Torrent) (*string, error) {
	// torrent file from s3 is downloaded to /tmp/current.torrent
	filename := filepath.Join(x.workingDir, torrent.Filename)
	file, err := os.Create(filename)
	if err != nil {
		x.logger.Errorf("failed to create file: %s", filename)
		return nil, err
	}
	defer func() { _ = file.Close() }()
	x.logger.Infof("starting downloading torrent from s3: %s", torrent.Filename)
	if err := x.s3Downloader.Download(file, env.BucketTorrentParsed.String(), torrent.Filename); err != nil {
		x.logger.Errorf("failed to download torrent from s3: %s", torrent.Filename)
		return nil, err
	}
	return &filename, nil
}

func (x *Processor) downloadTorrentContents(torrent *models.Torrent, timeout time.Duration) error {
	var lastProgress float64
	interval := time.Second * 5
	if err := x.torrentDownloader.Start(); err != nil {
		return err
	}
	// clean up resources for next execution
	defer func() {
		_ = x.torrentDownloader.Stop()
	}()
	if err := x.torrentDownloader.AddTorrent(x.workingDir, torrent.Filename,
		func(line *model.TorrentLine) {
			if line.Percent != lastProgress {
				// TODO: publish to torrent-info
				lastProgress = line.Percent
				x.logger.Info("download ", line.String())
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

func (*Processor) validateTorrent(torrent *models.Torrent) error {
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
