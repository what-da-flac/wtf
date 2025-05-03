package rest

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/what-da-flac/wtf/go-common/ifaces"
	"github.com/what-da-flac/wtf/services/gateway/internal/environment"
	"github.com/what-da-flac/wtf/services/gateway/internal/interfaces"
)

type Server struct {
	db              *sql.DB
	config          *environment.Config
	storePathFinder interfaces.PathFinder
	tempPathFinder  interfaces.PathFinder
	identifier      interfaces.Identifier
	logger          ifaces.Logger
	repository      interfaces.Repository
	timer           ifaces.Timer
}

func New(db *sql.DB, logger ifaces.Logger, repository interfaces.Repository) *Server {
	return &Server{
		db:         db,
		logger:     logger,
		repository: repository,
	}
}

func (x *Server) WithConfig(config *environment.Config) *Server {
	x.config = config
	return x
}

func (x *Server) WithPathFinders(storeFinder, tempFinder interfaces.PathFinder) *Server {
	x.storePathFinder = storeFinder
	x.tempPathFinder = tempFinder
	return x
}

func (x *Server) WithIdentifier(identifier ifaces.Identifier) *Server {
	x.identifier = identifier
	return x
}

func (x *Server) WithTimer(timer ifaces.Timer) *Server {
	x.timer = timer
	return x
}

func (x *Server) context(r *http.Request) context.Context {
	return r.Context()
}
