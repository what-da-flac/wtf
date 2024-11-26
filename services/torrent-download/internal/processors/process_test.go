package processors

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/what-da-flac/wtf/go-common/env"
	"github.com/what-da-flac/wtf/go-common/loggers"
	"github.com/what-da-flac/wtf/openapi/models"
	"github.com/what-da-flac/wtf/services/torrent-download/mocks"
)

func TestProcess_RemoveDownloadedFiles(t *testing.T) {
	var (
		removeAllCalled bool
		startCalled     bool
		stopCalled      bool
		filename        string
	)
	logger := loggers.MustNewDevelopmentLogger()
	torrentDownloader := &mocks.TorrentDownloaderMock{
		AddTorrentFunc: func(targetDir string, torrentFileName string) error {
			filename = torrentFileName
			return nil
		},
		RemoveTorrentsLeaveFilesFunc: func() error { return nil },
		RemoveTorrentsAndFilesFunc: func() error {
			removeAllCalled = true
			return nil
		},
		StartFunc: func() error {
			startCalled = true
			return nil
		},
		StopFunc: func() error {
			stopCalled = true
			return nil
		},
		WaitForDownloadFunc: func(wait time.Duration, interval time.Duration) bool {
			return false
		},
	}
	s3Downloader := &mocks.S3DownloaderMock{
		DownloadFunc: func(w io.WriterAt, bucket string, key string) error {
			assert.Equal(t, env.BucketTorrentParsed.String(), bucket)
			return nil
		},
	}
	torrent := &models.Torrent{
		Id:       "123",
		Filename: "123.torrent",
	}
	config := &env.Config{
		Volumes: env.Volumes{
			Downloads: env.Names(filepath.Join(os.TempDir(), "downloads")),
		},
	}
	config.Downloads.Timeout = time.Minute * 30
	elapsed, err := Process(logger, torrentDownloader, s3Downloader, torrent, config, os.TempDir())
	assert.Error(t, err)
	assert.Empty(t, elapsed)
	assert.True(t, removeAllCalled, "expected to call removeAll")
	assert.True(t, startCalled, "expected torrent downloader to start")
	assert.True(t, stopCalled, "expected torrent downloader to stop")
	assert.Equal(t, fmt.Errorf("torrent download timed out after: 30m0s"), err)
	assert.Equal(t, "123.torrent", filepath.Base(filename))
}
