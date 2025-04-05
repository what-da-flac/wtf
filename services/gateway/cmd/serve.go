package cmd

import (
	"net/http"

	"github.com/spf13/cobra"
	"github.com/what-da-flac/wtf/go-common/identifiers"
	"github.com/what-da-flac/wtf/go-common/pgpq"
	"github.com/what-da-flac/wtf/go-common/timers"
	"github.com/what-da-flac/wtf/openapi/gen/golang"
	"github.com/what-da-flac/wtf/services/gateway/internal/assets"
	"github.com/what-da-flac/wtf/services/gateway/internal/environment"
	"github.com/what-da-flac/wtf/services/gateway/internal/migrations"
	"github.com/what-da-flac/wtf/services/gateway/internal/rest"
	"go.uber.org/zap"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts service",
	RunE: func(cmd *cobra.Command, args []string) error {
		logger, err := zap.NewProductionConfig().Build()
		if err != nil {
			return err
		}
		return serve(logger, environment.New())
	},
}

func init() {
	cmd := serveCmd
	rootCmd.AddCommand(cmd)
}

func serve(zl *zap.Logger, config *environment.Config) error {
	logger := zl.Sugar()
	connStr := config.DB.URL
	logger.Infof("trying to connect to postgres at url: %s", config.DB.URL)
	db, err := pgpq.New(connStr)
	if err != nil {
		return err
	}
	defer func() { _ = db.Close() }()
	logger.Info("TODO: connect to rabbitmq")
	logger.Info("connected to db")
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
	api := rest.New(logger).
		WithConfig(config).
		WithTimer(timers.New()).
		WithIdentifier(identifier)
	mux := http.NewServeMux()
	handler := golang.HandlerFromMuxWithBaseURL(api, mux, apiURLPrefix)
	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           handler,
		ReadHeaderTimeout: config.HeaderTimeout,
	}
	defer func() { _ = srv.Close() }()

	logger.Infof("starting rest api at port: %s", port)
	return srv.ListenAndServe()
}

// applyHTTPMiddleware makes sure middlewares are executed in the order they were provided.
func applyHTTPMiddleware(h http.Handler, middlewares ...golang.MiddlewareFunc) http.Handler {
	chain := h
	for i := len(middlewares) - 1; i >= 0; i-- {
		m := middlewares[i]
		chain = m(chain)
	}
	return chain
}
