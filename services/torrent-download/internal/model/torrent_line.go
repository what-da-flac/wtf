package model

import (
	"fmt"
	"strconv"
	"strings"
)

// TorrentLine represents one torrent in the list
type TorrentLine struct {
	ID      string
	Eta     string
	Percent float64
}

func (x *TorrentLine) String() string {
	return fmt.Sprintf("progress: %.2f%% eta: %s",
		x.Percent*100,
		x.Eta)
}

func NewTorrentLine(line string) *TorrentLine {
	res := &TorrentLine{}
	values := strings.Fields(line)
	for i, v := range values {
		if i == 0 {
			idVal := strings.TrimSpace(v)
			if _, err := strconv.Atoi(idVal); err != nil {
				return nil
			}
		}
		switch i {
		case 0:
			res.ID = strings.TrimSpace(v)
		case 1:
			val, err := strconv.ParseFloat(strings.TrimSuffix(v, "%"), 64)
			if err != nil {
				continue
			}
			res.Percent = val / 100
		case 4:
			if values[3] == "Unknown" {
				res.Eta = values[3]
			} else {
				res.Eta = v + " " + values[i+1]
			}
		default:
			continue
		}
	}
	return res
}
