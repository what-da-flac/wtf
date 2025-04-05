package rest

import (
	"context"
	"net/http"

	"github.com/what-da-flac/wtf/go-common/ifaces"
	"github.com/what-da-flac/wtf/services/gateway/internal/environment"
	"github.com/what-da-flac/wtf/services/gateway/internal/interfaces"
)

type Server struct {
	config     *environment.Config
	identifier interfaces.Identifier
	logger     ifaces.Logger
	timer      ifaces.Timer
}

func New(logger ifaces.Logger) *Server {
	return &Server{
		logger: logger,
	}
}

func (x *Server) WithConfig(config *environment.Config) *Server {
	x.config = config
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
