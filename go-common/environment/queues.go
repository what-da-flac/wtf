package environment

type Queues struct {
	// MagnetParser a magnet_link in the payload
	MagnetParser string

	// TorrentDownload  a full torrent object
	TorrentDownload string

	// TorrentParser a torrent filename
	TorrentParser string

	// TorrentInfo a full torrent object
	TorrentInfo string
}

func newQueues() Queues {
	return Queues{
		MagnetParser:    "magnet-parser",
		TorrentDownload: "torrent-download",
		TorrentInfo:     "torrent-info",
		TorrentParser:   "torrent-parser",
	}
}
