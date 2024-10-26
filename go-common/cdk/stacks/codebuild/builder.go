package codebuild

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/jsii-runtime-go"
	"github.com/what-da-flac/wtf/go-common/cdk/stacks/common"
)

// Build builds a codebuild job based on EC2 engine.
func Build(stack awscdk.Stack, model *Model) {
	var (
		inLinePolicies  = make(map[string]awsiam.PolicyDocument)
		environments    = make(map[string]*awscodebuild.BuildEnvironmentVariable)
		managedPolicies []awsiam.IManagedPolicy
		webhookFilters  []awscodebuild.FilterGroup
	)
	if model.Source.Filter == WebhookTagRelease {
		// add default managed policies required to publish docker images to ecr
		managedPolicies = append(managedPolicies,
			awsiam.ManagedPolicy_FromAwsManagedPolicyName(
				jsii.String("AmazonEC2ContainerRegistryPowerUser"),
			),
		)
	}

	// set inline policies
	for k, v := range model.InlinePolicies {
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
	for _, v := range model.ManagedPolicies {
		managedPolicies = append(managedPolicies,
			awsiam.ManagedPolicy_FromAwsManagedPolicyName(jsii.String(v)))
	}
	// define role for codebuild job
	roleName := "CodeBuildRole_" + model.Name
	role := awsiam.NewRole(stack, jsii.String(roleName), &awsiam.RoleProps{
		AssumedBy:       awsiam.NewServicePrincipal(jsii.String("codebuild.amazonaws.com"), nil),
		ManagedPolicies: &managedPolicies,
		InlinePolicies:  &inLinePolicies,
	})
	// define environment variables
	for _, v := range model.Environments {
		environments[v.Name] = &awscodebuild.BuildEnvironmentVariable{
			Value: v.Value,
			Type:  v.Type.ToAWSCodebuild(),
		}
	}
	// define webhook filters
	switch model.Source.Filter {
	case WebhookPullRequest:
		webhookFilters = []awscodebuild.FilterGroup{
			awscodebuild.FilterGroup_InEventOf(awscodebuild.EventAction_PULL_REQUEST_CREATED),
			awscodebuild.FilterGroup_InEventOf(awscodebuild.EventAction_PULL_REQUEST_REOPENED),
			awscodebuild.FilterGroup_InEventOf(awscodebuild.EventAction_PULL_REQUEST_UPDATED),
		}
	case WebhookTagRelease:
		patternMatching := model.Source.PatternMatching
		if patternMatching == "" {
			patternMatching = "^refs/tags/.*"
		}
		webhookFilters = append(webhookFilters, awscodebuild.FilterGroup_InEventOf(awscodebuild.EventAction_PUSH).
			AndHeadRefIs(jsii.String(patternMatching)))
		// cannot build docker images without this setting
		model.Privileged = true
	}
	var dockerImage awscodebuild.IBuildImage
	switch model.Docker.RegistryType {
	case common.DockerRegistryCustom:
		dockerImage = awscodebuild.LinuxBuildImage_FromDockerRegistry(jsii.String(model.Docker.Url), &awscodebuild.DockerImageOptions{})
	case common.DockerRegistryAWS:
		dockerImage = awscodebuild.LinuxBuildImage_STANDARD_7_0()
	}
	// all codebuild jobs are assumed to be docker based, not lambdas
	awscodebuild.NewProject(
		stack,
		jsii.String(model.Name),
		&awscodebuild.ProjectProps{
			BuildSpec:   awscodebuild.BuildSpec_FromSourceFilename(jsii.String(model.Source.CodebuildScriptPath)),
			Description: jsii.String(model.Description),
			Environment: &awscodebuild.BuildEnvironment{
				BuildImage:           dockerImage,
				ComputeType:          model.ComputeTypeAWS,
				EnvironmentVariables: &environments,
				Privileged:           jsii.Bool(model.Privileged),
			},
			ProjectName: jsii.String(model.Name),
			Role:        role,
			Source: awscodebuild.Source_GitHub(
				&awscodebuild.GitHubSourceProps{
					Owner:             jsii.String(model.Source.Owner),
					Repo:              jsii.String(model.Source.Repo),
					ReportBuildStatus: jsii.Bool(true),
					Webhook:           jsii.Bool(true),
					WebhookFilters:    &webhookFilters,
				}),
		})
}
