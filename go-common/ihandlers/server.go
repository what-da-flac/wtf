package ihandlers

import (
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/what-da-flac/wtf/go-common/imodel"
)

var once sync.Once

// ServerHandler is a wrapper for grpc and http requests, so we can separate v1 and v2 routes.
type ServerHandler struct {
	// endpoints are the defined http endpoints with a route and optional http verbs
	endpoints []*imodel.Endpoint

	// handler is the handler that resolves routes
	handler http.Handler

	// internal middleware functions
	middlewares []func(http.Handler) http.Handler

	// strictSlash defines the trailing slash behavior for new routes.
	//
	// When true, if the route path is "/path/", accessing "/path" will perform a redirect
	// to the former and vice versa. In other words, your application will always
	// see the path as specified in the route.
	strictSlash bool
}

// NewServerHandler is the default constructor that routes all missing paths to a default handler.
func NewServerHandler() *ServerHandler {
	return &ServerHandler{}
}

// Build configures the endpoints and routers.
func (x *ServerHandler) Build() *ServerHandler {
	once.Do(func() {
		router := mux.NewRouter().
			StrictSlash(x.strictSlash)

		// all the missing endpoints will resolve the custom handler provided
		for _, ep := range x.endpoints {
			router.HandleFunc(ep.Path.String(), ep.Handler).Methods(ep.Verb)
		}
		x.handler = router
		x.handler = ApplyHTTPMiddleware(x.handler, x.middlewares...)
	})
	return x
}

func (x *ServerHandler) WithMiddlewares(middlewares ...func(http.Handler) http.Handler) *ServerHandler {
	x.middlewares = middlewares
	return x
}

func (x *ServerHandler) WithStrictSlash(strictSlash bool) *ServerHandler {
	x.strictSlash = strictSlash
	return x
}

func (x *ServerHandler) WithEndpoints(endpoints ...*imodel.Endpoint) *ServerHandler {
	x.endpoints = endpoints
	return x
}

func (x *ServerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if x.handler == nil {
		http.Error(w, "forgot to call build?", http.StatusServiceUnavailable)
		return
	}
	x.handler.ServeHTTP(w, r)
}
