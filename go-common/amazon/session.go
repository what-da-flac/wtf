package amazon

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/what-da-flac/wtf/go-common/environment"
)

type AWSSession struct {
	credential       *AWSCredential
	endpoint         *string
	s3ForcePathStyle *bool
	region           *string

	session *session.Session
}

func NewAWSSession() *AWSSession {
	return &AWSSession{}
}

func (x *AWSSession) WithCredential(credential *AWSCredential) *AWSSession {
	x.credential = credential
	return x
}

func (x *AWSSession) WithEndpoint(endpoint string) *AWSSession {
	x.endpoint = &endpoint
	return x
}

func (x *AWSSession) WithS3ForcePathStyle(forcePathStyle bool) *AWSSession {
	x.s3ForcePathStyle = &forcePathStyle
	return x
}

func (x *AWSSession) WithRegion(region string) *AWSSession {
	x.region = &region
	return x
}

func (x *AWSSession) Build() error {
	config := aws.NewConfig()
	if x.region != nil {
		config = config.WithRegion(*x.region)
	}
	if x.credential != nil {
		config = config.WithCredentials(x.credential.toAWSCred())
	}
	if x.endpoint != nil {
		config = config.WithEndpoint(*x.endpoint)
	}
	if x.s3ForcePathStyle != nil {
		config = config.WithS3ForcePathStyle(*x.s3ForcePathStyle)
	}
	sess, err := session.NewSession(config)
	if err != nil {
		return err
	}
	x.session = sess
	return nil
}

func (x *AWSSession) Session() *session.Session {
	return x.session
}

// NewAWSSessionFromEnvironment will read the environment variables and configure them accordingly.
// If no credentials are found, they will be used from the ecs task role.
// This function works also for local development as is.
func NewAWSSessionFromEnvironment() *AWSSession {
	c := environment.New().AWS
	return NewAWSSession().
		WithCredential(MustAWSCredentialFromConfiguration(c)).
		WithEndpoint(c.AWSEndpoint).
		WithRegion(c.AWSDefaultRegion).
		WithS3ForcePathStyle(c.AWSS3ForcePathStyle)
}
