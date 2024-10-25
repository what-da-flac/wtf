package lambda

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/what-da-flac/wtf/go-common/cdk/stacks/common"
)

func TestParse(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    []*Model
		wantErr bool
	}{
		{
			name: "happy path",
			args: args{
				filename: "test-data/lambda.yaml",
			},
			want: []*Model{
				{
					Code: Code{
						Docker: &common.Docker{
							RegistryType: common.DockerRegistryCustom,
							Url:          "torrent-download:v0.0.0.13",
						},
					},
					Environment: []common.Environment{
						{
							Name:  "BEER",
							Type:  common.EnvironmentTypeText,
							Value: "is good",
						},
					},
					EphemeralStorageSizeGb: aws.Float64(10),
					MemorySizeMb:           aws.Float64(512),
					Name:                   "torrent-download",
					TimeoutSeconds:         aws.Float64(900),
					Trigger: &Trigger{
						Type: TriggerTypeSQS,
					},
				},
				{
					Code: Code{
						Docker: &common.Docker{
							RegistryType: common.DockerRegistryCustom,
							Url:          "torrent-processing:v0.0.1",
						},
					},
					Environment: []common.Environment{
						{
							Name:  "WINE",
							Type:  common.EnvironmentTypeText,
							Value: "is better",
						},
					},
					EphemeralStorageSizeGb: aws.Float64(8),
					MemorySizeMb:           aws.Float64(128),
					Name:                   "torrent-processing",
					TimeoutSeconds:         aws.Float64(900),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			require.Len(t, got, len(tt.want))
			for i, got := range got {
				want := tt.want[i]
				assert.Equal(t, *got, *want)
			}
		})
	}
}
