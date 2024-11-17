package env

// defines names for each queue to be reused among all services.
const (
	QueueMagnetParser    Names = "magnet-parser"
	QueueTorrentDownload Names = "torrent-download"
	QueueTorrentInfo     Names = "torrent-info"
	QueueTorrentParser   Names = "torrent-parser"
)
