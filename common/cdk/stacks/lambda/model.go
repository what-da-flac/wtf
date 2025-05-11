package lambda

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/jsii-runtime-go"
	"github.com/what-da-flac/wtf/common/cdk/stacks/common"
)

type Model struct {
	Code                   Code                     `yaml:"Code"`
	Environment            []common.Environment     `yaml:"Environment"`
	EphemeralStorageSizeGb *float64                 `yaml:"EphemeralStorageSizeGb"`
	InlinePolicies         map[string]common.Policy `yaml:"InlinePolicies"`
	ManagedPolicies        []string                 `yaml:"ManagedPolicies"`
	MemorySizeMb           *float64                 `yaml:"MemorySizeMb"`
	Name                   string                   `yaml:"Name"`
	TimeoutSeconds         *float64                 `yaml:"TimeoutSeconds"`
	Trigger                *Trigger                 `yaml:"Trigger"`
}

func (x *Model) Validate() error {
	if x.Code.Docker == nil {
		return fmt.Errorf("code from docker must be specified")
	}
	if x.Code.Docker.RegistryType != common.DockerRegistryCustom {
		return fmt.Errorf("code from docker must use custom registry")
	}
	return nil
}

func (x *Model) ECRArn() string {
	return "arn:aws:ecr:" + common.AWSRegion() + ":" + common.AWSAccount() + ":repository/" + x.Name
}

func (x *Model) defaultPolicy() awsiam.PolicyDocument {
	return awsiam.NewPolicyDocument(
		&awsiam.PolicyDocumentProps{
			Statements: &[]awsiam.PolicyStatement{
				awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
					Actions: &[]*string{
						jsii.String("logs:CreateLogGroup"),
						jsii.String("logs:CreateLogStream"),
						jsii.String("logs:PutLogEvents"),
					},
					Effect:    awsiam.Effect_ALLOW,
					Resources: &[]*string{jsii.String("*")},
				}),
			},
		},
	)
}

func (x *Model) Role(stack awscdk.Stack) awsiam.Role {
	var (
		inLinePolicies  = make(map[string]awsiam.PolicyDocument)
		managedPolicies []awsiam.IManagedPolicy
	)
	// set default inline policies for all lambdas
	inLinePolicies["_default"] = x.defaultPolicy()
	// set inline policies
	for k, v := range x.InlinePolicies {
		var resources []*string
		for _, resource := range v.Resources {
			resources = append(resources, jsii.String(resource))
		}
		inLinePolicies[k] = awsiam.NewPolicyDocument(&awsiam.PolicyDocumentProps{
			Statements: &[]awsiam.PolicyStatement{
				awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
					Actions:   &[]*string{jsii.String(v.Action)},
					Effect:    awsiam.Effect_ALLOW,
					Resources: &resources,
				}),
			},
		})
	}
	// set managed policies
	for _, v := range x.ManagedPolicies {
		managedPolicies = append(managedPolicies,
			awsiam.ManagedPolicy_FromAwsManagedPolicyName(jsii.String(v)))
	}
	roleName := "LambdaRole_" + x.Name
	return awsiam.NewRole(stack, jsii.String(roleName), &awsiam.RoleProps{
		AssumedBy:       awsiam.NewServicePrincipal(jsii.String("lambda.amazonaws.com"), nil),
		ManagedPolicies: &managedPolicies,
		InlinePolicies:  &inLinePolicies,
	})
}
