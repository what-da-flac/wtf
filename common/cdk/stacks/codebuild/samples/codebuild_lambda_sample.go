package samples

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/jsii-runtime-go"
)

// TODO: remove this file, kept only as reference.
func CodebuildLambda(stack awscdk.Stack) {
	role := awsiam.NewRole(stack, jsii.String("CodeBuildRole"), &awsiam.RoleProps{
		AssumedBy: awsiam.NewServicePrincipal(jsii.String("codebuild.amazonaws.com"), nil),
		ManagedPolicies: &[]awsiam.IManagedPolicy{
			awsiam.ManagedPolicy_FromAwsManagedPolicyName(jsii.String("AmazonS3FullAccess")),
		},
		InlinePolicies: &map[string]awsiam.PolicyDocument{
			"s3": awsiam.NewPolicyDocument(&awsiam.PolicyDocumentProps{
				Statements: &[]awsiam.PolicyStatement{
					awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
						Actions:   &[]*string{jsii.String("secretsmanager:GetSecretValue")},
						Effect:    awsiam.Effect_ALLOW,
						Resources: &[]*string{jsii.String("arn:aws:secretsmanager:*:*")},
					}),
				},
			}),
		},
	})
	awscodebuild.NewProject(
		stack,
		jsii.String("wtf-ui-release"),
		&awscodebuild.ProjectProps{
			BuildSpec: awscodebuild.BuildSpec_FromObject(&map[string]interface{}{
				"version": "0.2",
				"phases": map[string]interface{}{
					"install": map[string]interface{}{
						"commands": []string{
							"npm install",
							"npm run build",
						},
					},
				},
			}),
			Description: jsii.String("Builds wtf ui and sends files to s3 bucket"),
			Environment: &awscodebuild.BuildEnvironment{
				BuildImage:  awscodebuild.LinuxArmLambdaBuildImage_AMAZON_LINUX_2023_NODE_20(),
				ComputeType: awscodebuild.ComputeType_LAMBDA_4GB,
				EnvironmentVariables: &map[string]*awscodebuild.BuildEnvironmentVariable{
					"REACT_APP_BASE_API_URL": {
						Value: "https://dev.app.publicitux.com/api",
						Type:  awscodebuild.BuildEnvironmentVariableType_PLAINTEXT,
					},
					"REACT_APP_GOOGLE_CLIENT_ID": {
						Value: "google-client-id",
						Type:  awscodebuild.BuildEnvironmentVariableType_SECRETS_MANAGER,
					},
					"REACT_APP_GOOGLE_CLIENT_SECRET": {
						Value: "google-client-secret",
						Type:  awscodebuild.BuildEnvironmentVariableType_SECRETS_MANAGER,
					},
					"REACT_APP_GOOGLE_API_KEY": {
						Value: "google-api-key",
						Type:  awscodebuild.BuildEnvironmentVariableType_SECRETS_MANAGER,
					},
				},
			},
			ProjectName: jsii.String("wtf-ui-release"),
			Role:        role,
			Source: awscodebuild.Source_GitHub(
				&awscodebuild.GitHubSourceProps{
					Owner:             jsii.String("tech-component"),
					Repo:              jsii.String("wtf-ui"),
					ReportBuildStatus: jsii.Bool(true),
					Webhook:           jsii.Bool(true),
					WebhookFilters: &[]awscodebuild.FilterGroup{
						awscodebuild.FilterGroup_InEventOf(awscodebuild.EventAction_PUSH),
					},
				}),
		})
}
