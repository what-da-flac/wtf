package cmd

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_nextVersion(t *testing.T) {
	type args struct {
		r           io.Reader
		versionType string
	}
	tests := []struct {
		name    string
		args    args
		wantW   string
		wantErr bool
	}{
		{
			name: "happy path lambdas",
			args: args{
				r: bytes.NewBufferString(`cdk.0.0.1
cdk.0.0.2
cdk.0.0.3
cdk.0.0.4
cdk.0.0.5
cdk.0.0.6
cdk.0.0.7
cdk.0.0.8
cdk.0.0.9
docker.0.0.1
docker.0.0.2
docker.0.0.3
docker.0.0.4
gateway.0.0.1
lambda.0.0.1
lambda.0.0.10
lambda.0.0.11
lambda.0.0.12
lambda.0.0.2
lambda.0.0.3
lambda.0.0.4
lambda.0.0.5
lambda.0.0.6
lambda.0.0.7
lambda.0.0.8
lambda.0.0.9
ui.0.0.1
v0.0.1
v0.0.2
v0.0.3
v0.0.4
`),
				versionType: "lambda",
			},
			wantW:   "lambda.0.0.13\n",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			err := nextVersion(tt.args.r, w, tt.args.versionType)
			if (err != nil) != tt.wantErr {
				t.Errorf("nextVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.wantW, w.String())
		})
	}
}
