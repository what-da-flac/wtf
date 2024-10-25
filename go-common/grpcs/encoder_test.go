package grpcs

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/what-da-flac/wtf/go-common/protos/gen/go/protos/common_domain"
	"google.golang.org/protobuf/proto"
)

func TestProtobufEncoder(t *testing.T) {
	type args struct {
		m proto.Message
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "happy path",
			args: args{
				m: &common_domain.User{
					Id:        "1",
					Email:     "test@example.com",
					ImageUrl:  "https://something.com/image",
					IsDeleted: true,
				},
			},
			want:    "CgExQhB0ZXN0QGV4YW1wbGUuY29tShtodHRwczovL3NvbWV0aGluZy5jb20vaW1hZ2VQAQ==",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ProtobufEncoder(tt.args.m)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			if err != nil {
				return
			}
			assert.Equal(t, tt.want, got)
			m := &common_domain.User{}
			err = ProtobufDecoder(got, m)
			assert.NoError(t, err)
			assert.True(t, proto.Equal(tt.args.m, m))
		})
	}
}
