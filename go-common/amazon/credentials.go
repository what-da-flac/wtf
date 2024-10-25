package amazon

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/what-da-flac/wtf/go-common/environment"
)

type AWSCredential struct {
	accessKey string
	secret    string
	token     string
}

func NewAWSCredential(accessKey, secret, token string) *AWSCredential {
	return &AWSCredential{
		accessKey: accessKey,
		secret:    secret,
		token:     token,
	}
}

func (x *AWSCredential) toAWSCred() *credentials.Credentials {
	return credentials.NewStaticCredentials(x.accessKey, x.secret, x.token)
}

func MustAWSCredentialFromConfiguration(aws environment.AWS) *AWSCredential {
	if !aws.HasCredentials() {
		panic("not aws credentials found AWS_ACCESS_KEY_ID / AWS_SECRET_ACCESS_KEY / AWS_DEFAULT_REGION")
	}
	return NewAWSCredential(aws.AWSAccessKeyId, aws.AWSSecretAccessKey, aws.AWSSessionToken)
}
