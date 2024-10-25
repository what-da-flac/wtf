package s3

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/what-da-flac/wtf/go-common/cdk/stacks/common"
)

func TestParse(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    *Model
		wantErr bool
	}{
		{
			name: "happy-path.yaml",
			args: args{
				filename: "happy-path.yaml",
			},
			want: &Model{
				AutoDeleteObjects: true,
				EnforceSSL:        false,
				InlinePolicies: map[string]common.Policy{
					"s3": {
						Action: "s3:*",
						Resources: []string{
							"arn:aws:s3:::wtf-ui.what-da-flac.com",
							"arn:aws:s3:::wtf-ui.what-da-flac.com/*",
						},
					},
				},
				Name:                 "wtf-ui.what-da-flac.com",
				BlockPublicAccess:    false,
				Versioned:            false,
				WebsiteIndexDocument: "index.html",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filename := filepath.Join("test-data", tt.args.filename)
			got, err := Parse(filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("fromYAML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
