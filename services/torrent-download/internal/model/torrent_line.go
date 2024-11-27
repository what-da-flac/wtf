package model

import "fmt"

// TorrentLine represents one torrent in the list
type TorrentLine struct {
	ID      string
	Percent float64
}

func (x *TorrentLine) String() string {
	return fmt.Sprintf("progress: %.2f%%", x.Percent*100)
}
