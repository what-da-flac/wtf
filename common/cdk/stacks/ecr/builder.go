package ecr

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/jsii-runtime-go"
)

func Build(stack awscdk.Stack, base *Model) {
	var (
		mutability    awsecr.TagMutability
		removalPolicy awscdk.RemovalPolicy
	)
	if base.Mutable {
		mutability = awsecr.TagMutability_MUTABLE
	} else {
		mutability = awsecr.TagMutability_IMMUTABLE
	}
	if base.RemoveOnDestroy {
		removalPolicy = awscdk.RemovalPolicy_DESTROY
	} else {
		removalPolicy = awscdk.RemovalPolicy_RETAIN
	}
	awsecr.NewRepository(stack, jsii.String(base.Name), &awsecr.RepositoryProps{
		EmptyOnDelete:      jsii.Bool(base.EmptyOnDelete),
		RemovalPolicy:      removalPolicy,
		RepositoryName:     jsii.String(base.Name),
		ImageTagMutability: mutability,
	})
}
