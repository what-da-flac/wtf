package s3

import (
	"strconv"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/jsii-runtime-go"
)

func Build(stack awscdk.Stack, base *Model) {
	var rules []*awss3.LifecycleRule
	if val := base.ExpirationDays; val > 0 {
		rules = append(rules, &awss3.LifecycleRule{
			Enabled:    jsii.Bool(true),
			Expiration: awscdk.Duration_Days(jsii.Number(val)),
			Id:         jsii.String(base.Name + "-expires-" + strconv.Itoa(val) + "-days"),
		},
		)
	}
	props := &awss3.BucketProps{
		BucketName:           jsii.String(base.Name),
		EnforceSSL:           jsii.Bool(base.EnforceSSL),
		LifecycleRules:       &rules,
		PublicReadAccess:     jsii.Bool(base.BlockPublicAccess),
		Versioned:            jsii.Bool(base.Versioned),
		WebsiteIndexDocument: jsii.String(base.WebsiteIndexDocument),
	}
	if base.AutoDeleteObjects && base.RemoveOnDestroy {
		props.AutoDeleteObjects = jsii.Bool(true)
		props.RemovalPolicy = awscdk.RemovalPolicy_DESTROY
	}
	if base.BlockPublicAccess {
		props.BlockPublicAccess = awss3.BlockPublicAccess_BLOCK_ALL()
	} else {
		props.BlockPublicAccess = awss3.BlockPublicAccess_BLOCK_ACLS()
	}
	awss3.NewBucket(stack, jsii.String(base.Name), props)
}
