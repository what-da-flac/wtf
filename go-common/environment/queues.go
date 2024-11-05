package environment

type Queues struct {
	MagnetParser  string
	TorrentParser string
}

func newQueues() Queues {
	return Queues{
		MagnetParser:  "magnet-parser",
		TorrentParser: "torrent-parser",
	}
}
