package ihandlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

type PathParam string

// NewPathParam looks up in the request path, for variable values.
func NewPathParam(r *http.Request, k string) PathParam {
	return PathParam(mux.Vars(r)[k])
}

func (x PathParam) String() string {
	return string(x)
}
