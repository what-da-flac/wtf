package rest

import (
	"context"
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/what-da-flac/wtf/gateway/internal/environment"
	"github.com/what-da-flac/wtf/gateway/internal/interfaces"
	"github.com/what-da-flac/wtf/go-common/ihandlers"
	"github.com/what-da-flac/wtf/openapi/models"
)

type Server struct {
	config        *environment.Config
	identifier    interfaces.Identifier
	repository    interfaces.Repository
	timer         interfaces.Timer
	messageSender interfaces.MessageSender
}

func New() *Server {
	return &Server{}
}

func (x *Server) WithConfig(config *environment.Config) *Server {
	x.config = config
	return x
}

func (x *Server) WithIdentifier(identifier interfaces.Identifier) *Server {
	x.identifier = identifier
	return x
}

func (x *Server) WithRepository(repository interfaces.Repository) *Server {
	x.repository = repository
	return x
}

func (x *Server) WithTimer(timer interfaces.Timer) *Server {
	x.timer = timer
	return x
}

func (x *Server) WithMessageSender(messageSender interfaces.MessageSender) *Server {
	x.messageSender = messageSender
	return x
}

func (x *Server) context(r *http.Request) context.Context {
	return r.Context()
}

func (x *Server) ReadUserFromContext(ctx context.Context) *models.User {
	res := &models.User{}
	if val := ihandlers.UserFromContext(ctx); val != nil {
		if err := copier.Copy(res, val); err != nil {
			return nil
		}
		return res
	}
	return nil
}
