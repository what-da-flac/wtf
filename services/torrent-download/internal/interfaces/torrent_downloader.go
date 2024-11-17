package interfaces

import "time"

//go:generate moq -out ../../mocks/torrent_downloader.go -pkg mocks . TorrentDownloader
type TorrentDownloader interface {

	// AddTorrent appends a torrent to the queue and starts downloading its files.
	AddTorrent(targetDir, torrentFileName string) error

	// ClearTorrents clears torrents but keeps files.
	ClearTorrents() error

	// RemoveAll clears torrents and files.
	RemoveAll() error

	// Start runs background processes.
	Start() error

	// Stop kills background processes.
	Stop() error

	// WaitForDownload returns true if download was completed.
	WaitForDownload(wait, interval time.Duration) bool
}
