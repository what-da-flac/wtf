package restful

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/what-da-flac/wtf/openapi/gen/golang"
)

type ReadCloser struct {
	data *bytes.Buffer
}

func (r *ReadCloser) Read(p []byte) (n int, err error) {
	return r.data.Read(p)
}

func (r *ReadCloser) Close() error {
	return nil
}

func TestReadRequest(t *testing.T) {
	type args struct {
		r *http.Request
	}
	type testCase[T any] struct {
		name    string
		args    args
		want    *T
		wantErr bool
	}
	tests := []testCase[golang.AudioFile]{
		{
			name: "happy path",
			args: args{
				r: &http.Request{
					Method: http.MethodGet,
					Header: http.Header{
						"Content-Type": []string{contentTypeJSON},
					},
					Body: &ReadCloser{data: bytes.NewBufferString(`
{
  "album": "2112",
  "duration_seconds": 99,
  "title": "Something for Nothing"
}
`)},
				},
			},
			want: &golang.AudioFile{
				Album:           "2112",
				DurationSeconds: 99,
				Title:           "Something for Nothing",
			},
			wantErr: false,
		},
		{
			name: "invalid content type",
			args: args{
				r: &http.Request{
					Method: http.MethodGet,
					Header: http.Header{},
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadRequest[golang.AudioFile](tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
