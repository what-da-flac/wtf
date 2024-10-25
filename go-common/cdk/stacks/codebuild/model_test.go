package codebuild

import (
	"path/filepath"
	"testing"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild"
	"github.com/stretchr/testify/assert"
	"github.com/what-da-flac/wtf/go-common/cdk/stacks/common"
)

func TestFromYAML(t *testing.T) {
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
			name: "releaser.yaml",
			args: args{
				filename: "releaser.yaml",
			},
			want: &Model{
				ComputeTypeAWS: awscodebuild.ComputeType_MEDIUM,
				Description:    "Builds wtf ui and sends files to s3 bucket",
				Environments: []common.Environment{
					{
						Name:  "GITHUB_TOKEN",
						Type:  common.EnvironmentTypeSecret,
						Value: "github-token",
					},
					{
						Name:  "SERVICE_NAME",
						Type:  common.EnvironmentTypeText,
						Value: "wtf",
					},
				},
				InlinePolicies: map[string]common.Policy{
					"s3": {
						Action: "s3:*",
						Resources: []string{
							"arn:aws:s3:::wtf-ui.what-da-flac.com",
							"arn:aws:s3:::wtf-ui.what-da-flac.com/*",
						},
					},
				},
				ManagedPolicies: []string{
					"S3Admin",
				},
				Name: "test-job",
				Source: Github{
					CodebuildScriptPath: "/codebuild/test.yaml",
					Repo:                "wtf-devops",
					Owner:               "tech-component",
					Filter:              WebhookPullRequest,
					PatternMatching:     "^refs/tags/docker.*",
				},
			},
			wantErr: false,
		},
		{
			name: "releaser-invalid-environment.yaml",
			args: args{
				filename: "releaser-invalid-environment.yaml",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "releaser-invalid-filter.yaml",
			args: args{
				filename: "releaser-invalid-filter.yaml",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filename := filepath.Join("test-data", tt.args.filename)
			got, err := Parse(filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("BaseFromYAML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
