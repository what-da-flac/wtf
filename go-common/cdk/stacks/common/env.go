package common

import "os"

func AWSAccount() string {
	return os.Getenv("CDK_DEFAULT_ACCOUNT")
}

func AWSRegion() string {
	return os.Getenv("CDK_DEFAULT_REGION")
}
