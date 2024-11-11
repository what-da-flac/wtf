package env

type BucketName string

func (x BucketName) String() string { return string(x) }

const (
	BucketTorrentParsed BucketName = "wtf.torrent-parsed"
)
