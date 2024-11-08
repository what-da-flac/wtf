package env

type Buckets struct {
	TorrentParsed string
}

func newBuckets() Buckets {
	return Buckets{
		TorrentParsed: "wtf.torrent-parsed",
	}
}
