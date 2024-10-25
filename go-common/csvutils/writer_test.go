package csvutils

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrite(t *testing.T) {
	type args struct {
		headers []string
		rowFn   func(index int) ([]string, error)
	}
	tests := []struct {
		name    string
		args    args
		wantW   string
		wantErr bool
	}{
		{
			name: "happy path",
			args: args{
				headers: []string{
					"FirstName",
					"lastName",
					"Age",
				},
				rowFn: func(index int) ([]string, error) {
					switch index {
					case 0:
						return []string{
							"John",
							"Doe",
							"36",
						}, nil
					case 1:
						return []string{
							"Deborah",
							"Harry",
							"77",
						}, nil
					default:
						return nil, io.EOF
					}
				},
			},
			wantW:   "FirstName,lastName,Age\nJohn,Doe,36\nDeborah,Harry,77\n",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			err := Write(w, tt.args.headers, tt.args.rowFn)
			if (err != nil) != tt.wantErr {
				t.Errorf("want: %v got: %v", tt.wantErr, err)
			}
			assert.EqualValues(t, tt.wantW, w.String())
		})
	}
}
