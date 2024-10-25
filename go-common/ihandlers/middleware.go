package ihandlers

import "net/http"

// ApplyHTTPMiddleware attaches the middlewares to a handler, preserving array elements order.
func ApplyHTTPMiddleware(h http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	chain := h
	for i := len(middlewares) - 1; i >= 0; i-- {
		m := middlewares[i]
		chain = m(chain)
	}
	return chain
}
