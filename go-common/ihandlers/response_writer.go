package ihandlers

import "net/http"

// ResponseWriter is a custom implementation of a response writer, that stores the status code,
// so it can be used in a middleware scenario.
type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
	err        error
}

func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		ResponseWriter: w,
		statusCode:     http.StatusOK,
		err:            nil,
	}
}

func (w *ResponseWriter) WriteHeader(code int) {
	// ignore assigning default status code
	if code == http.StatusOK {
		return
	}
	// status code can only be assigned once
	if w.statusCode != 0 && w.statusCode != http.StatusOK {
		return
	}
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *ResponseWriter) Write(b []byte) (int, error) {
	if w.err != nil {
		return 0, w.err
	}
	return w.ResponseWriter.Write(b)
}

func (w *ResponseWriter) IsError() bool {
	return w.statusCode >= 400
}

func (w *ResponseWriter) StatusCode() int { return w.statusCode }
