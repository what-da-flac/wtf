package i_am

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/jsii-runtime-go"
)

type GroupUsers struct {
	Group *ModelGroup
	Users []*ModelUser
}

func Build(stack awscdk.Stack, groups []*ModelGroup, users []*ModelUser) {
	groupUsers := linkGroupsToUsers(groups, users)
	for _, gu := range groupUsers {
		build(stack, gu.Group, gu.Users)
	}
}

// linkGroupsToUsers creates structure to link groups with users
// all groups and all users are present in the result
// and user/group membership is considered.
func linkGroupsToUsers(groups []*ModelGroup, users []*ModelUser) []*GroupUsers {
	var res []*GroupUsers
	// create one entry for each group, no matter the user membership
	for _, group := range groups {
		item := &GroupUsers{
			Group: group,
		}
		res = append(res, item)
		// attach users to each group
		for _, user := range users {
			var ok bool
			for _, g := range user.Groups {
				if g == group.Name {
					item.Users = append(item.Users, user)
					ok = true
					break
				}
			}
			if !ok {
				break
			}
		}
	}
	// find users without a group membership, and include in the results with a nil group
	var missingUsers []*ModelUser
	for _, user := range users {
		var ok bool
		for _, gu := range res {
			for _, g := range gu.Users {
				if g == user {
					ok = true
					break
				}
			}
			if ok {
				break
			}
		}
		if !ok {
			missingUsers = append(missingUsers, user)
		}
	}
	if len(missingUsers) > 0 {
		res = append(res, &GroupUsers{
			Users: missingUsers,
		})
	}
	return res
}

func build(stack awscdk.Stack, g *ModelGroup, users []*ModelUser) {
	// consider the case for users without group membership
	if g == nil {
		for _, user := range users {
			awsiam.NewUser(stack, jsii.String(user.Username), &awsiam.UserProps{
				UserName: jsii.String(user.Username),
			})
		}
		return
	}

	group := awsiam.NewGroup(stack, jsii.String(g.Name), &awsiam.GroupProps{
		GroupName: jsii.String(g.Name),
	})
	for _, policy := range g.Policies {
		var resources []*string
		for _, res := range policy.Resources {
			resources = append(resources, jsii.String(res))
		}
		group.AttachInlinePolicy(awsiam.NewPolicy(stack, jsii.String(policy.Name), &awsiam.PolicyProps{
			Statements: &[]awsiam.PolicyStatement{
				awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
					Actions: &[]*string{
						jsii.String(policy.Action),
					},
					Effect:    awsiam.Effect_ALLOW,
					Resources: &resources,
				}),
			},
		}))
	}
	for _, managedPolicy := range g.ManagedPolicies {
		group.AddManagedPolicy(
			awsiam.ManagedPolicy_FromAwsManagedPolicyName(
				jsii.String(managedPolicy),
			),
		)
	}
	for _, u := range users {
		user := awsiam.NewUser(stack, jsii.String(u.Username), &awsiam.UserProps{
			UserName: jsii.String(u.Username),
		})
		group.AddUser(user)
	}
}
