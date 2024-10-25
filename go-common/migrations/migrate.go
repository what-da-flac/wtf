package migrations

import (
	"errors"
	"io/fs"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

// MigrateFS executes a db migration against a set of files.
// Driver is assumed to be Postgres.
//
// example:
//
//	files := assets.Files()
//	if err := assets.MigratePG(files, "files/migrations", connStr); err != nil {
//		logger.Errorf("running db migrations: %s", err)
//		return err
//	}
//	logger.Info("db migrations completed successfully")
func MigrateFS(files fs.FS, filePath, dbURL string) error {
	const sourceName = "iofs"
	driver, err := iofs.New(files, filePath)
	if err != nil {
		return err
	}
	m, err := migrate.NewWithSourceInstance(sourceName, driver, dbURL)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return err
		}
	}
	return nil
}

// MigrateDirectory will execute migrations on provided dir path,
// for each of the files it contains.
//
// example:
// migrations.MigrateDirectory("postgres://postgres:password@localhost:5432?sslmode=disable", "file://internal/assets/files/migrations")
func MigrateDirectory(connStr, dir string) error {
	m, err := migrate.New(dir, connStr)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return err
		}
	}
	return nil
}
