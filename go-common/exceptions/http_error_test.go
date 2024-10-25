package exceptions

import (
	"database/sql"
	"net/http"
	"reflect"
	"testing"
)

func TestNewHTTPError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want *HTTPError
	}{
		{
			name: "nil error",
			args: args{
				err: nil,
			},
			want: &HTTPError{},
		},
		{
			name: "from httpError",
			args: args{
				err: NewHTTPError(nil).WithCustomMessage("forbidden").WithStatusCode(http.StatusForbidden),
			},
			want: &HTTPError{
				Code:    http.StatusForbidden,
				Message: "forbidden",
			},
		},
		{
			name: "sql not found 404",
			args: args{
				err: sql.ErrNoRows,
			},
			want: &HTTPError{
				Code:    http.StatusNotFound,
				Message: "not found",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHTTPError(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHTTPError() = %v, want %v", got, tt.want)
			}
		})
	}
}
