package pgrepo

import (
	"context"

	"github.com/what-da-flac/wtf/openapi/models"
)

func (x *PgRepo) InsertTorrent(_ context.Context, torrent *models.Torrent) error {
	db := x.GORM()
	dto := torrentFromModel(torrent)
	return db.Create(dto).Error
}

func (x *PgRepo) UpdateTorrent(ctx context.Context, torrent *models.Torrent) error {
	db := x.GORM()
	dto := torrentFromModel(torrent)
	return db.Save(dto).Error
}

func (x *PgRepo) SelectTorrent(_ context.Context, id string) (*models.Torrent, error) {
	db := x.GORM()
	row := &TorrentDto{}
	if err := db.Model(row).Where("id = ?", id).First(row).Error; err != nil {
		// Handle other types of database errors
		return nil, err
	}
	return row.toModel(), nil
}

func (x *PgRepo) ListTorrents(ctx context.Context, params models.GetV1TorrentsParams) ([]*models.Torrent, error) {
	var (
		res  []*models.Torrent
		rows []*TorrentDto
	)
	db := x.GORM()
	db = db.Model(&TorrentDto{}).
		Limit(params.Limit).
		Order(params.SortField + " " + params.SortDirection)
	if val := params.Offset; val != nil {
		db = db.Offset(*val)
	}
	if val := params.Status; val != nil {
		db = db.Where("status = ?", *val)
	}
	db = db.Find(&rows)
	if err := db.Error; err != nil {
		return nil, err
	}
	for _, row := range rows {
		res = append(res, row.toModel())
	}
	return res, nil
}

func (x *PgRepo) ListTorrentStatuses(ctx context.Context) []string {
	return []string{
		string(models.Downloaded),
		string(models.Downloading),
		string(models.Queued),
		string(models.Parsed),
		string(models.Pending),
	}
}

func (x *PgRepo) InsertTorrentFile(ctx context.Context, file *models.TorrentFile) error {
	db := x.GORM()
	return db.Create(torrentFileFromModel(file)).Error
}

func (x *PgRepo) DeleteTorrentFiles(ctx context.Context, torrentId string) error {
	db := x.GORM()
	return db.Where("torrent_id = ?", torrentId).Delete(&TorrentFileDto{}).Error
}

func (x *PgRepo) SelectTorrentFiles(ctx context.Context, id string) ([]*models.TorrentFile, error) {
	var (
		res  []*models.TorrentFile
		rows []*TorrentFileDto
	)
	db := x.GORM()
	db = db.Model(&TorrentFileDto{}).Where("torrent_id = ?", id)
	db = db.Find(&rows)
	if err := db.Error; err != nil {
		return nil, err
	}
	for _, row := range rows {
		res = append(res, row.toModel())
	}
	return res, nil
}
