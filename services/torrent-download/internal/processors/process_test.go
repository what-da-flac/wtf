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
	"github.com/what-da-flac/wtf/services/torrent-download/internal/interfaces"
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
		AddTorrentFunc: func(targetDir, torrentFileName string, fn interfaces.ProgressFn) error {
			filename = torrentFileName
			return nil
		},
		ClearTorrentsFunc: func() error { return nil },
		RemoveAllFunc: func() error {
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
		DownloadFunc: func(w io.Writer, bucket string, key string) error {
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
	publisher := &mocks.PublisherMock{
		PublishFunc: func(data []byte) error {
			t.Log("publishing: ", string(data))
			return nil
		},
	}
	config.Downloads.Timeout = time.Minute * 30
	x := NewProcessor(logger, torrentDownloader, s3Downloader, publisher)
	elapsed, err := x.Process(torrent, config, os.TempDir())
	assert.Error(t, err)
	assert.Empty(t, elapsed)
	assert.True(t, removeAllCalled, "expected to call removeAll")
	assert.False(t, startCalled, "not expected torrent downloader to start")
	assert.False(t, stopCalled, "not expected torrent downloader to stop")
	assert.Equal(t, fmt.Errorf("torrent download timed out after: 30m0s"), err)
	assert.Equal(t, "123.torrent", filepath.Base(filename))
}
