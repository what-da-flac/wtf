package i_am

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/what-da-flac/wtf/go-common/cdk/stacks/common"
)

func TestParseGroups(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    ModelGroups
		wantErr bool
	}{
		{
			name: "happy path",
			args: args{
				filename: "test-data/groups.yaml",
			},
			want: []*ModelGroup{
				{
					Name: "test-group-1",
					Policies: []*common.Policy{
						{
							Action: "s3:*",
							Name:   "s3-full",
							Resources: []string{
								"arn:aws:s3:::wtf-ui.what-da-flac.com",
								"arn:aws:s3:::wtf-ui.what-da-flac.com/*",
							},
						},
						{
							Action: "ecr:*",
							Name:   "ecs-full",
							Resources: []string{
								"*",
							},
						},
					},
					ManagedPolicies: []string{
						"AmazonS3ReadOnlyAccess",
					},
				},
				{
					Name: "test-group-2",
					Policies: []*common.Policy{
						{
							Action: "ecs:*",
							Name:   "ecs-full",
							Resources: []string{
								"arn:aws:ecs:::*",
							},
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseGroups(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("BaseFromYAML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.ElementsMatch(t, tt.want.UnPtr(), got.UnPtr())
		})
	}
}

func TestParseUsers(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    []*ModelUser
		wantErr bool
	}{
		{
			name: "happy path",
			args: args{
				filename: "test-data/users.yaml",
			},
			want: []*ModelUser{
				{
					Username: "mauleyzaola",
					Groups: []string{
						"administrators",
						"power_users",
						"users",
					},
				},
				{
					Username: "test@example.com",
					Groups: []string{
						"users",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseUsers(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseUsers() got = %v, want %v", got, tt.want)
			}
		})
	}
}
