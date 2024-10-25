package ihandlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestNewPathParam(t *testing.T) {
	type args struct {
		request *http.Request
		k       string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "happy path",
			args: args{
				request: func() *http.Request {
					req := httptest.NewRequest(http.MethodGet, "/", nil)
					req = mux.SetURLVars(req, map[string]string{
						"my-id": "abc-123",
					})
					return req
				}(),
				k: "my-id",
			},
			want: "abc-123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewPathParam(tt.args.request, tt.args.k)
			assert.Equal(t, tt.want, got.String())
		})
	}
}
