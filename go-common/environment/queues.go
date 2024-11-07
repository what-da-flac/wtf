package environment

type Queues struct {
	// MagnetParser downloads a torrent file from magnet link
	MagnetParser string

	// TorrentDownload downloads torrent content
	TorrentDownload string

	// TorrentInfo gets info from torrent file
	TorrentInfo string

	// TorrentParser receives torrent information and stores in db
	TorrentParser string
}

func newQueues() Queues {
	return Queues{
		MagnetParser:    "magnet-parser",
		TorrentDownload: "torrent-download",
		TorrentInfo:     "torrent-info",
		TorrentParser:   "torrent-parser",
	}
}
