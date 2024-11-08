package env

import "github.com/spf13/viper"

//nolint:gosec
const (
	// these are the standard environment variable names from AWS documentation,
	// and we should stick to them, if decide to use them.
	// Although it is a better practice to inherit permissions from the service role instead.
	// https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-envvars.html
	envVarAWSAccessKeyId = "AWS_ACCESS_KEY_ID"

	envVarAWSSecretAccessKey = "AWS_SECRET_ACCESS_KEY"
	envVarAWSSessionToken    = "AWS_SESSION_TOKEN"
	envVarAWSDefaultRegion   = "AWS_DEFAULT_REGION"

	// AWS_ENDPOINT is used for localstack
	// https://docs.localstack.cloud/references/network-troubleshooting/endpoint-url/
	envVarAWSEndpoint      = "AWS_ENDPOINT"
	envVarS3ForcePathStyle = "S3_FORCE_PATH_STYLE"
)

type AWS struct {
	AWSAccessKeyId      string
	AWSSecretAccessKey  string
	AWSSessionToken     string
	AWSDefaultRegion    string
	AWSEndpoint         string
	AWSS3ForcePathStyle bool
}

func newAWS() AWS {
	const defaultAWSRegion = "us-east-2"
	aws := AWS{
		AWSAccessKeyId:      viper.GetString(envVarAWSAccessKeyId),
		AWSSecretAccessKey:  viper.GetString(envVarAWSSecretAccessKey),
		AWSSessionToken:     viper.GetString(envVarAWSSessionToken),
		AWSDefaultRegion:    viper.GetString(envVarAWSDefaultRegion),
		AWSEndpoint:         viper.GetString(envVarAWSEndpoint),
		AWSS3ForcePathStyle: viper.GetBool(envVarS3ForcePathStyle),
	}
	if aws.AWSDefaultRegion == "" {
		aws.AWSDefaultRegion = defaultAWSRegion
	}
	return aws
}

func (x *AWS) HasCredentials() bool {
	return x.AWSAccessKeyId != "" && x.AWSSecretAccessKey != ""
}
