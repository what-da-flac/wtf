package grpcs

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/what-da-flac/wtf/go-common/protos/gen/go/protos/common_domain"
	"google.golang.org/protobuf/proto"
)

func TestToJSON(t *testing.T) {
	type args struct {
		m                 proto.Message
		useCamelCase      bool
		includeNullValues bool
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "user snake case",
			args: args{
				m: &common_domain.User{
					Id:        "1",
					Email:     "test@example.com",
					ImageUrl:  "https://some-url.com/1234.png",
					IsDeleted: false,
				},
				useCamelCase:      false,
				includeNullValues: true,
			},
			want:    "{\"id\":\"1\", \"email\":\"test@example.com\", \"image_url\":\"https://some-url.com/1234.png\", \"is_deleted\":false}",
			wantErr: false,
		},
		{
			name: "user snake case omit nulls",
			args: args{
				m: &common_domain.User{
					Id:        "1",
					Email:     "test@example.com",
					ImageUrl:  "https://some-url.com/1234.png",
					IsDeleted: false,
				},
				useCamelCase:      false,
				includeNullValues: false,
			},
			want:    "{\"id\":\"1\", \"email\":\"test@example.com\", \"image_url\":\"https://some-url.com/1234.png\"}",
			wantErr: false,
		},
		{
			name: "user camel case",
			args: args{
				m: &common_domain.User{
					Id:        "1",
					Email:     "test@example.com",
					ImageUrl:  "https://some-url.com/1234.png",
					IsDeleted: false,
				},
				useCamelCase:      true,
				includeNullValues: true,
			},
			want:    "{\"id\":\"1\", \"email\":\"test@example.com\", \"imageUrl\":\"https://some-url.com/1234.png\", \"isDeleted\":false}",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToJSON(tt.args.m, tt.args.useCamelCase, tt.args.includeNullValues)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			k1, err := extractKeys([]byte(tt.want))
			assert.NoError(t, err)
			k2, err := extractKeys(got)
			assert.NoError(t, err)
			// compare keys match and ignore the rest (we only care about how field names are written)
			assert.Equal(t, len(k1), len(k2))
			for k := range k1 {
				delete(k2, k)
			}
			assert.Empty(t, k2)
		})
	}
}

func extractKeys(v []byte) (map[string]interface{}, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(v, &m); err != nil {
		return nil, err
	}
	return m, nil
}
