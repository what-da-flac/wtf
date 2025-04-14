package main

import (
	"net/http"

	"github.com/what-da-flac/wtf/services/gateway/internal/storages"

	_ "github.com/lib/pq"
	"github.com/what-da-flac/wtf/go-common/identifiers"
	"github.com/what-da-flac/wtf/go-common/pgpq"
	"github.com/what-da-flac/wtf/go-common/repositories"
	"github.com/what-da-flac/wtf/go-common/timers"
	"github.com/what-da-flac/wtf/openapi/gen/golang"
	"github.com/what-da-flac/wtf/services/gateway/internal/assets"
	"github.com/what-da-flac/wtf/services/gateway/internal/environment"
	"github.com/what-da-flac/wtf/services/gateway/internal/migrations"
	"github.com/what-da-flac/wtf/services/gateway/internal/rest"
	"go.uber.org/zap"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	logger, err := zap.NewProductionConfig().Build()
	if err != nil {
		return err
	}
	return serve(logger)
}

func serve(zl *zap.Logger) error {
	config := environment.New()
	connStr := config.DB.URL
	db, err := pgpq.New(connStr)
	if err != nil {
		return err
	}
	defer func() { _ = db.Close() }()
	repository, err := repositories.NewPgRepo(db, connStr, false)
	if err != nil {
		return err
	}

	logger := zl.Sugar()
	if err = migrations.MigrateFS(
		assets.MigrationFiles(),
		"files/migrations",
		config.DB.URL,
	); err != nil {
		return err
	}
	logger.Info("db migrations applied successfully")

	port := config.Port
	apiURLPrefix := config.APIUrlPrefix
	identifier := identifiers.NewIdentifier()
	fileStorage := storages.NewLocal(config.PathToSaveFiles)
	api := rest.New(db, logger, repository).
		WithConfig(config).
		WithFileStorage(fileStorage).
		WithIdentifier(identifier).
		WithTimer(timers.New())
	mux := http.NewServeMux()
	handler := golang.HandlerFromMuxWithBaseURL(api, mux, apiURLPrefix)
	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           handler,
		ReadHeaderTimeout: config.HeaderTimeout,
	}
	defer func() { _ = srv.Close() }()

	logger.Infof("serving from %s:%s", config.APIUrlPrefix, config.Port)
	return srv.ListenAndServe()
}
