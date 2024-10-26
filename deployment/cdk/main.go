package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewStack(scope constructs.Construct, id string, props *awscdk.StackProps) awscdk.Stack {
	return awscdk.NewStack(scope, &id, props)
}

func main() {
	defer jsii.Close()
	app := awscdk.NewApp(nil)
	if err := run(app); err != nil {
		panic(err)
	}
	app.Synth(nil)
}

func run(app awscdk.App) error {
	// runs stacks defined in this directory as YAML files
	stackProps, err := ReadStacks(".")
	if err != nil {
		return err
	}

	// apply all stack props
	for _, s := range stackProps {
		st := NewStack(app, s.Name, &awscdk.StackProps{
			Env:         env(),
			StackName:   jsii.String(s.Name),
			Description: jsii.String(s.Description),
		})
		Deploy(st, s.BaseDir)
	}

	return nil
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	return &awscdk.Environment{
		Region: jsii.String("us-east-2"),
	}
}
