package amazon

import (
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"golang.org/x/net/context"
)

type AWSSession struct {
	region *string

	cfg *aws.Config
}

func NewAWSSession() *AWSSession {
	return &AWSSession{}
}

func (x *AWSSession) WithRegion(region string) *AWSSession {
	x.region = &region
	return x
}

func (x *AWSSession) Build() error {
	var options []func(*config.LoadOptions) error
	if x.region != nil {
		options = append(options, config.WithRegion(*x.region))
	}
	cfg, err := config.LoadDefaultConfig(context.TODO(), options...)
	if err != nil {
		return err
	}
	x.cfg = &cfg
	return nil
}

func (x *AWSSession) Session() *aws.Config {
	return x.cfg
}

// NewAWSSessionFromEnvironment will read the environment variables and configure them accordingly.
// If no credentials are found, they will be used from the ecs task role.
// This function works also for local development as is.
func NewAWSSessionFromEnvironment() *AWSSession {
	return NewAWSSession().WithRegion(os.Getenv("AWS_DEFAULT_REGION"))
}
