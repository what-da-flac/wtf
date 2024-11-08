package env

type QueueName string

func (x QueueName) String() string { return string(x) }

// defines names for each queue to be reused among all services.
const (
	QueueMagnetParser    QueueName = "magnet-parser"
	QueueTorrentDownload QueueName = "torrent-download"
	QueueTorrentInfo     QueueName = "torrent-info"
	QueueTorrentParser   QueueName = "torrent-parser"
)
