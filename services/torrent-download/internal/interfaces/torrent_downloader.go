package interfaces

import "time"

//go:generate moq -out ../../mocks/torrent_downloader.go -pkg mocks . TorrentDownloader
type TorrentDownloader interface {

	// AddTorrent appends a torrent to the queue and starts downloading its files.
	AddTorrent(targetDir, torrentFileName string) error

	// RemoveTorrentsLeaveFiles clears torrents but keeps files.
	RemoveTorrentsLeaveFiles() error

	// RemoveTorrentsAndFiles clears torrents and files.
	RemoveTorrentsAndFiles() error

	// Start runs background processes.
	Start() error

	// Stop kills background processes.
	Stop() error

	// WaitForDownload returns true if download was completed.
	WaitForDownload(wait, interval time.Duration) bool
}
