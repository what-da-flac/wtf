package sqs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/jsii-runtime-go"
)

func Build(stack awscdk.Stack, base *Model) {
	var removalPolicy awscdk.RemovalPolicy
	if base.RemoveOnDestroy {
		removalPolicy = awscdk.RemovalPolicy_DESTROY
	} else {
		removalPolicy = awscdk.RemovalPolicy_RETAIN
	}
	awssqs.NewQueue(stack, jsii.String(base.Name), &awssqs.QueueProps{
		DeliveryDelay:     awscdk.Duration_Seconds(jsii.Number(base.DeliveryDelaySeconds)),
		Encryption:        awssqs.QueueEncryption_KMS_MANAGED,
		EnforceSSL:        jsii.Bool(true),
		Fifo:              jsii.Bool(true),
		QueueName:         jsii.String(base.Name),
		RemovalPolicy:     removalPolicy,
		RetentionPeriod:   awscdk.Duration_Seconds(jsii.Number(base.RetentionPeriodSeconds)),
		VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(base.VisibilityTimeoutSeconds)),
	})
}
