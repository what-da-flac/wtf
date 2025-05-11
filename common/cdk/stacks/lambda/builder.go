package lambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambdaeventsources"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/jsii-runtime-go"
)

func Build(stack awscdk.Stack, m *Model) {
	var (
		ephemeralStorageSize *float64
		environment          = make(map[string]*string)
		memorySize           *float64
		timeout              awscdk.Duration
	)
	if val := m.EphemeralStorageSizeGb; val != nil {
		ephemeralStorageSize = val
	}
	if val := m.MemorySizeMb; val != nil {
		memorySize = val
	}
	if val := m.TimeoutSeconds; val != nil {
		timeout = awscdk.Duration_Seconds(val)
	} else {
		// default is 15 seconds in aws
		timeout = awscdk.Duration_Seconds(jsii.Number(15))
	}
	// define environment variables
	for _, v := range m.Environment {
		environment[v.Name] = jsii.String(v.Value)
	}
	repoName, repoTag := m.Code.Docker.RepositoryTagName()
	ecrRepo := awsecr.Repository_FromRepositoryAttributes(stack,
		jsii.String("ecr-"+m.Name),
		&awsecr.RepositoryAttributes{
			RepositoryArn:  jsii.String(m.ECRArn()),
			RepositoryName: jsii.String(repoName),
		})
	function := awslambda.NewFunction(stack,
		jsii.String("lambda-"+m.Name),
		&awslambda.FunctionProps{
			EphemeralStorageSize: awscdk.Size_Gibibytes(ephemeralStorageSize),
			FunctionName:         jsii.String(m.Name),
			MemorySize:           memorySize,
			Role:                 m.Role(stack),
			Timeout:              timeout,
			Code: awslambda.Code_FromEcrImage(ecrRepo, &awslambda.EcrImageCodeProps{
				TagOrDigest: jsii.String(repoTag),
			}),
			Handler:     awslambda.Handler_FROM_IMAGE(),
			Runtime:     awslambda.Runtime_FROM_IMAGE(),
			Environment: &environment,
		})
	if m.Trigger == nil {
		return
	}
	//nolint:gocritic
	switch m.Trigger.Type {
	case TriggerTypeSQS:
		var queueName string
		// TODO:allow to choose fifo or standard
		queueName = m.Name + ".fifo"
		queue := awssqs.NewQueue(stack, jsii.String(queueName), &awssqs.QueueProps{
			Fifo:              jsii.Bool(true),
			QueueName:         jsii.String(queueName),
			RemovalPolicy:     awscdk.RemovalPolicy_DESTROY,
			RetentionPeriod:   awscdk.Duration_Seconds(jsii.Number(1209600)),
			VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(10800)),
		})
		// convention: we create sqs queue with the same name as lambda receiver
		function.AddEventSource(
			awslambdaeventsources.NewSqsEventSource(queue,
				&awslambdaeventsources.SqsEventSourceProps{
					// TODO: allow to optionally set these fields
					//BatchSize:      jsii.Number(1),
					Enabled: jsii.Bool(true),
					//MaxConcurrency: jsii.Number(1),
				},
			),
		)
	}
}
