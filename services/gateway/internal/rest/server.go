package rest

import (
	"context"
	"net/http"

	"github.com/what-da-flac/wtf/services/gateway/internal/environment"
	interfaces2 "github.com/what-da-flac/wtf/services/gateway/internal/interfaces"

	"github.com/jinzhu/copier"
	"github.com/what-da-flac/wtf/go-common/ihandlers"
	"github.com/what-da-flac/wtf/openapi/models"
)

type Server struct {
	config        *environment.Config
	identifier    interfaces2.Identifier
	repository    interfaces2.Repository
	timer         interfaces2.Timer
	messageSender interfaces2.MessageSender
}

func New() *Server {
	return &Server{}
}

func (x *Server) WithConfig(config *environment.Config) *Server {
	x.config = config
	return x
}

func (x *Server) WithIdentifier(identifier interfaces2.Identifier) *Server {
	x.identifier = identifier
	return x
}

func (x *Server) WithRepository(repository interfaces2.Repository) *Server {
	x.repository = repository
	return x
}

func (x *Server) WithTimer(timer interfaces2.Timer) *Server {
	x.timer = timer
	return x
}

func (x *Server) WithMessageSender(messageSender interfaces2.MessageSender) *Server {
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
