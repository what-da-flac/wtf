package model

import "fmt"

// TorrentLine represents one torrent in the list
type TorrentLine struct {
	ID         string
	Downloaded float64
	Eta        string
	Percent    float64
}

func (x *TorrentLine) String() string {
	return fmt.Sprintf("downloaded Mb: %.1f progress: %.2f%% eta: %s",
		x.Downloaded/1024/1024,
		x.Percent*100,
		x.Eta)
}
