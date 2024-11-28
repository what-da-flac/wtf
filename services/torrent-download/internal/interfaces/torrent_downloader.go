package interfaces

import (
	"time"

	"github.com/what-da-flac/wtf/services/torrent-download/internal/model"
)

//go:generate moq -out ../../mocks/torrent_downloader.go -pkg mocks . TorrentDownloader
type TorrentDownloader interface {

	// AddTorrent appends a torrent to the queue and starts downloading its files.
	AddTorrent(targetDir, torrentFileName string, progressFn ProgressFn) error

	// ClearTorrents removes torrents but keeps its files.
	ClearTorrents() error

	// RemoveAll deletes torrents and files.
	RemoveAll() error

	// Start runs background processes.
	Start() error

	// Stop kills background processes.
	Stop() error

	// WaitForDownload returns true if download was completed.
	WaitForDownload(wait, interval time.Duration) bool
}

type ProgressFn func(line *model.TorrentLine)
