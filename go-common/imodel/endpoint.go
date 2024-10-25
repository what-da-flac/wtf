package imodel

import (
	"net/http"
	"path/filepath"
)

// Endpoint represents a specific endpoint that matches a unique path/verb.
type Endpoint struct {
	Handler http.HandlerFunc
	Verb    string
	Path    EndpointPath
}

// NewEndpoint is the default constructor for an endpoint
// if no method is provided, all http verbs will be available.
func NewEndpoint(prefix, path string, method string, h http.HandlerFunc) *Endpoint {
	return &Endpoint{
		Handler: h,
		Verb:    method,
		Path:    EndpointPath(filepath.Join(prefix, path)),
	}
}
