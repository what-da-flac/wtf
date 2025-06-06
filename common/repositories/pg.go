package repositories

import (
	"database/sql"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PgRepo struct {
	db           *sql.DB
	_db          *gorm.DB
	defaultLimit int
}

func NewPgRepo(db *sql.DB, uri string, debug bool) (*PgRepo, error) {
	_db, err := NewGormConn(uri, debug)
	if err != nil {
		return nil, err
	}
	return &PgRepo{
		defaultLimit: 10,
		_db:          _db,
		db:           db,
	}, nil
}

func (x *PgRepo) GORM() *gorm.DB { return x._db }

func NewGormConn(uri string, debug bool) (*gorm.DB, error) {
	var lg logger.Interface
	if debug {
		lg = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Info,
				IgnoreRecordNotFoundError: true,
				ParameterizedQueries:      false,
				Colorful:                  true,
			},
		)
	}
	config := &gorm.Config{
		Logger: lg,
	}
	return gorm.Open(postgres.Open(uri), config)
}
