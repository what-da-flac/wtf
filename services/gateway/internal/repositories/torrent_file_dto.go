package repositories

import (
	"github.com/jinzhu/copier"
	"github.com/what-da-flac/wtf/openapi/models"
)

type TorrentFileDto struct {
	Id        string
	FileName  string
	FileSize  string
	TorrentId string
}

func (x *TorrentFileDto) TableName() string { return "torrent_file" }

func (x *TorrentFileDto) toModel() *models.TorrentFile {
	res := &models.TorrentFile{}
	if err := copier.Copy(res, x); err != nil {
		return nil
	}
	return res
}

func torrentFileFromModel(r *models.TorrentFile) *TorrentFileDto {
	res := &TorrentFileDto{}
	if err := copier.Copy(res, r); err != nil {
		return nil
	}
	return res
}
