package repositories

import (
	"database/sql"

	"github.com/what-da-flac/wtf/go-common/ifaces"
	"gorm.io/gorm"
)

type PG struct {
	db  ifaces.DB
	_db *gorm.DB

	defaultLimit int
}

func NewPG(db *sql.DB, uri string) (*PG, error) {
	const defaultRowLimit = 10
	_db, err := NewGormConn(uri)
	if err != nil {
		return nil, err
	}
	return &PG{
		_db:          _db,
		db:           db,
		defaultLimit: defaultRowLimit,
	}, nil
}

func (x *PG) GORM() *gorm.DB { return x._db }
