package codebuild

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild"
	"github.com/what-da-flac/wtf/go-common/cdk/stacks/common"
)

// Model defines codebuild job properties.
type Model struct {
	ComputeType     string                   `yaml:"ComputeType"`
	ComputeTypeAWS  awscodebuild.ComputeType `yaml:"-"`
	Description     string                   `yaml:"Description"`
	Docker          common.Docker            `yaml:"Docker"`
	Environments    []common.Environment     `yaml:"Environments"`
	InlinePolicies  map[string]common.Policy `yaml:"InlinePolicies"`
	ManagedPolicies []string                 `yaml:"ManagedPolicies"`
	Name            string                   `yaml:"Name"`
	Privileged      bool                     `yaml:"Privileged"`
	Source          Github                   `yaml:"Source"`
}

func (x *Model) Validate() error {
	var ct awscodebuild.ComputeType
	switch x.ComputeType {
	case "SMALL":
		ct = awscodebuild.ComputeType_SMALL
	case "MEDIUM":
		ct = awscodebuild.ComputeType_MEDIUM
	case "LARGE":
		ct = awscodebuild.ComputeType_LARGE
	//	stop supporting huge images, too much money
	// case "X_LARGE":
	//	ct = awscodebuild.ComputeType_X_LARGE
	// case "X2_LARGE":
	//	ct = awscodebuild.ComputeType_X2_LARGE
	default:
		return fmt.Errorf("unsupported model type: %s", x.ComputeType)
	}
	x.ComputeType = ""
	x.ComputeTypeAWS = ct
	return nil
}
