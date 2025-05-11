package i_am

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_linkGroupsToUsers(t *testing.T) {
	type args struct {
		groups []*ModelGroup
		users  []*ModelUser
	}
	tests := []struct {
		name string
		args args
		want []*GroupUsers
	}{
		{
			name: "happy path",
			args: args{
				groups: []*ModelGroup{
					{
						Name: "administrators",
					},
					{
						Name: "standalone",
						ManagedPolicies: []string{
							"AmazonS3ReadOnlyAccess",
						},
					},
				},
				users: []*ModelUser{
					{
						Username: "u1",
						Groups: []string{
							"administrators",
							"non-existing",
						},
					},
					{
						Username: "u2",
						Groups: []string{
							"administrators",
						},
					},
					{
						Username: "no-group-user",
					},
				},
			},
			want: []*GroupUsers{
				{
					Group: &ModelGroup{
						Name: "administrators",
					},
					Users: []*ModelUser{
						{
							Username: "u1",
							Groups: []string{
								"administrators",
								"non-existing",
							},
						},
						{
							Username: "u2",
							Groups: []string{
								"administrators",
							},
						},
					},
				},
				{
					Group: &ModelGroup{
						Name: "standalone",
						ManagedPolicies: []string{
							"AmazonS3ReadOnlyAccess",
						},
					},
				},
				{
					Group: nil,
					Users: []*ModelUser{
						{
							Username: "no-group-user",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := linkGroupsToUsers(tt.args.groups, tt.args.users)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}
