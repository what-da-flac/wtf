package rest

import (
	"context"
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/what-da-flac/wtf/go-common/env"
	"github.com/what-da-flac/wtf/go-common/ifaces"
	"github.com/what-da-flac/wtf/go-common/ihandlers"
	"github.com/what-da-flac/wtf/go-common/rabbits"
	"github.com/what-da-flac/wtf/openapi/gen/golang"
	"github.com/what-da-flac/wtf/services/gateway/internal/environment"
	interfaces2 "github.com/what-da-flac/wtf/services/gateway/internal/interfaces"
)

type Server struct {
	config     *environment.Config
	identifier interfaces2.Identifier
	logger     ifaces.Logger
	repository interfaces2.Repository
	timer      ifaces.Timer
	publishers map[env.Names]ifaces.Publisher

	rabbitURL string
}

func New(logger ifaces.Logger, rabbitURL string) *Server {
	return &Server{
		logger:     logger,
		publishers: make(map[env.Names]ifaces.Publisher),
		rabbitURL:  rabbitURL,
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

func (x *Server) WithRepository(repository interfaces2.Repository) *Server {
	x.repository = repository
	return x
}

func (x *Server) WithTimer(timer ifaces.Timer) *Server {
	x.timer = timer
	return x
}

func (x *Server) AddPublisher(key env.Names) *Server {
	p := rabbits.NewPublisher(x.logger, key, x.rabbitURL)
	if err := p.Build(); err != nil {
		panic(err)
	}
	x.publishers[key] = p
	return x
}

func (x *Server) publisher(key env.Names) ifaces.Publisher {
	return x.publishers[key]
}

func (x *Server) context(r *http.Request) context.Context {
	return r.Context()
}

func (x *Server) ReadUserFromContext(ctx context.Context) *golang.User {
	res := &golang.User{}
	if val := ihandlers.UserFromContext(ctx); val != nil {
		if err := copier.Copy(res, val); err != nil {
			return nil
		}
		return res
	}
	return nil
}
