package ihandlers

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplyHTTPMiddleware(t *testing.T) {

	firstFunc := func(w io.Writer) func(http.Handler) http.Handler {
		return func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				_, _ = w.Write([]byte("1"))
				next.ServeHTTP(writer, request)
			})
		}
	}

	secondFunc := func(w io.Writer) func(http.Handler) http.Handler {
		return func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				_, _ = w.Write([]byte("2"))
				next.ServeHTTP(writer, request)
			})
		}
	}

	thirdFunc := func(w io.Writer) func(http.Handler) http.Handler {
		return func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				_, _ = w.Write([]byte("3"))
				next.ServeHTTP(writer, request)
			})
		}
	}

	buffer := &bytes.Buffer{}
	testHandler := http.HandlerFunc(func(writer http.ResponseWriter, r *http.Request) {
		if _, err := io.Copy(buffer, r.Body); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	})
	handler := ApplyHTTPMiddleware(testHandler, firstFunc(buffer), secondFunc(buffer), thirdFunc(buffer))
	req := httptest.NewRequest(http.MethodGet, "https://example.com/test", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	assert.Equal(t, "123", buffer.String())
}
