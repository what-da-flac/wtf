package cmd

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/what-da-flac/wtf/services/gateway/internal/assets"
	"github.com/what-da-flac/wtf/services/gateway/internal/environment"
	"github.com/what-da-flac/wtf/services/gateway/internal/interfaces"
	"github.com/what-da-flac/wtf/services/gateway/internal/repositories"
	"github.com/what-da-flac/wtf/services/gateway/internal/rest"
	"github.com/what-da-flac/wtf/services/gateway/internal/senders"
	stores2 "github.com/what-da-flac/wtf/services/gateway/internal/stores"

	"github.com/jinzhu/copier"
	"github.com/spf13/cobra"
	"github.com/what-da-flac/wtf/go-common/amazon"
	"github.com/what-da-flac/wtf/go-common/identifiers"
	"github.com/what-da-flac/wtf/go-common/ihandlers"
	"github.com/what-da-flac/wtf/go-common/imodel"
	"github.com/what-da-flac/wtf/go-common/migrations"
	"github.com/what-da-flac/wtf/go-common/pgpq"
	"github.com/what-da-flac/wtf/go-common/sso"
	"github.com/what-da-flac/wtf/go-common/timers"
	"github.com/what-da-flac/wtf/openapi/models"
	"go.uber.org/zap"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
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
	db, err := pgpq.New(connStr)
	if err != nil {
		return err
	}
	defer func() { _ = db.Close() }()
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
	repository, err := repositories.NewPG(db, connStr)
	if err != nil {
		return err
	}
	awsSession := amazon.NewAWSSessionFromEnvironment()
	if err = awsSession.Build(); err != nil {
		return err
	}
	sess := awsSession.Session()
	identifier := identifiers.NewIdentifier()
	messageSender := senders.NewMessageSender(sess, logger, identifier)
	api := rest.New().
		WithConfig(config).
		WithTimer(timers.New()).
		WithIdentifier(identifier).
		WithRepository(repository).
		WithMessageSender(messageSender)
	mux := http.NewServeMux()
	handler := models.HandlerFromMuxWithBaseURL(api, mux, apiURLPrefix)
	middlewares, err := configureMiddlewares(config, repository)
	if err != nil {
		return err
	}
	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           applyHTTPMiddleware(handler, middlewares...),
		ReadHeaderTimeout: config.HeaderTimeout,
	}
	defer func() { _ = srv.Close() }()

	logger.Infof("starting rest api at port: %s", port)
	return srv.ListenAndServe()
}

func configureMiddlewares(config *environment.Config, repository interfaces.Repository) ([]models.MiddlewareFunc, error) {
	var res []models.MiddlewareFunc
	// TODO: use a cache mechanism
	roleResolver := stores2.NewRoleStore(repository)
	userResolver := stores2.NewUserStore()
	tokenValidator := sso.NewGoogleValidator()
	// read permissions from assets
	localAssets := assets.Files()
	permissionsFs, err := localAssets.Open(filepath.Join("files", "endpoints", "permissions.yaml"))
	if err != nil {
		return nil, err
	}
	endpointPermissions, err := imodel.ParsePermissions(permissionsFs)
	if err != nil {
		return nil, err
	}
	convertUserFn := func(user any) (*imodel.User, error) {
		googleUser, ok := user.(*imodel.GoogleUserInfo)
		if !ok {
			return nil, fmt.Errorf("cannot cast user to *model.GoogleUserInfo")
		}
		u := googleUser.ToUser()
		internalUser := &imodel.User{}
		if err := copier.Copy(internalUser, u); err != nil {
			return nil, err
		}
		return u, nil
	}
	cors := ihandlers.CORSMiddleware(
		ihandlers.DefaultCORSMaxAge, ihandlers.DefaultCORSHeaders,
		ihandlers.DefaultCORSMethods, ihandlers.DefaultCORSOrigins)
	instantiateFn := func() any { return &imodel.GoogleUserInfo{} }
	jwtMiddleware := ihandlers.JWTMiddleware(
		config.APIUrlPrefix, config.HeaderTimeout,
		endpointPermissions, userResolver, tokenValidator,
		instantiateFn, convertUserFn,
	)
	roleMiddleware := ihandlers.RoleMiddleware(config.APIUrlPrefix, endpointPermissions, roleResolver.Roles)
	userMiddleware := ihandlers.UserMiddleware(config.APIUrlPrefix, endpointPermissions, userResolver)
	res = append(res, cors, jwtMiddleware, roleMiddleware, userMiddleware)
	return res, nil
}

// applyHTTPMiddleware makes sure middlewares are executed in the order they were provided.
func applyHTTPMiddleware(h http.Handler, middlewares ...models.MiddlewareFunc) http.Handler {
	chain := h
	for i := len(middlewares) - 1; i >= 0; i-- {
		m := middlewares[i]
		chain = m(chain)
	}
	return chain
}
