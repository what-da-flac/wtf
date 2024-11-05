package domain

import (
	"github.com/what-da-flac/wtf/openapi/models"
)

// Torrent represents the metadata of a torrent file.
type Torrent struct {
	Name       string    `json:"name"`
	PieceCount int       `json:"piece_count"`
	PieceSize  string    `json:"piece_size"`
	TotalSize  string    `json:"total_size"`
	Privacy    string    `json:"privacy"`
	Trackers   []Tracker `json:"trackers"`
	Files      []File    `json:"files"`
}

// Tracker represents the tracker details of a torrent.
type Tracker struct {
	Tier int    `json:"tier"`
	URL  string `json:"url"`
}

// File represents the files included in a torrent.
type File struct {
	FileName string `json:"file_name"`
	FileSize string `json:"file_size"`
}

func (x *Torrent) ToDomain() *models.Torrent {
	res := &models.Torrent{
		Name:       x.Name,
		PieceCount: x.PieceCount,
		PieceSize:  x.PieceSize,
		Privacy:    x.Privacy,
		TotalSize:  x.TotalSize,
	}
	for _, file := range x.Files {
		res.Files = append(res.Files, models.TorrentFile{
			FileName: file.FileName,
			FileSize: file.FileSize,
		})
	}
	return res
}
