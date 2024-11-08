package pgrepo

import (
	"database/sql"
	"time"

	"github.com/jinzhu/copier"
	"github.com/what-da-flac/wtf/openapi/models"
)

type TorrentDto struct {
	Created    time.Time
	Filename   string
	Hash       string
	Id         string
	LastError  sql.NullString
	MagnetLink string
	Name       string
	PieceCount int
	PieceSize  string
	Privacy    string
	Status     string
	TotalSize  string
	Updated    *time.Time
	UserId     string
}

func (x *TorrentDto) TableName() string { return "torrent" }

func (x *TorrentDto) toModel() *models.Torrent {
	res := &models.Torrent{}
	if err := copier.Copy(res, x); err != nil {
		return nil
	}
	if val := x.UserId; val != "" {
		res.User = &models.User{
			Id: val,
		}
	}
	return res
}

func torrentFromModel(r *models.Torrent) *TorrentDto {
	res := &TorrentDto{}
	if err := copier.Copy(res, r); err != nil {
		return nil
	}
	if val := r.User; val != nil {
		res.UserId = val.Id
	}
	res.LastError.Valid = len(r.LastError) != 0
	res.LastError.String = r.LastError
	return res
}
