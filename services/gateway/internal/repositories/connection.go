package repositories

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormConn(uri string) (*gorm.DB, error) {
	config := &gorm.Config{
		// uncomment below to get generated sql in logs
		/*
			Logger: logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
				logger.Config{
					SlowThreshold:             time.Second, // Slow SQL threshold
					LogLevel:                  logger.Info, // Log level
					IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
					ParameterizedQueries:      true,        // Don't include params in the SQL log
					Colorful:                  false,       // Disable color
				},
			),
		*/
	}
	return gorm.Open(postgres.Open(uri), config)
}
