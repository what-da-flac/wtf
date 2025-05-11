package i_am

import "github.com/aws/aws-cdk-go/awscdk/v2"

func Iam(stack awscdk.Stack, groups []ModelGroup) {
	iamGroups(stack, groups)
}

func iamGroups(stack awscdk.Stack, groups []ModelGroup) {}
