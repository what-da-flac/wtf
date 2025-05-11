package common

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild"
)

type EnvironmentType int

const (
	EnvironmentTypeText   EnvironmentType = iota
	EnvironmentTypeSecret EnvironmentType = iota
)

func (x EnvironmentType) ToAWSCodebuild() awscodebuild.BuildEnvironmentVariableType {
	switch x {
	case EnvironmentTypeText:
		return awscodebuild.BuildEnvironmentVariableType_PLAINTEXT
	case EnvironmentTypeSecret:
		return awscodebuild.BuildEnvironmentVariableType_SECRETS_MANAGER
	default:
		return awscodebuild.BuildEnvironmentVariableType_PLAINTEXT
	}
}

type Environment struct {
	Name       string          `yaml:"Name"`
	Type       EnvironmentType `yaml:"-"`
	TypeString string          `yaml:"Type"`
	Value      string          `yaml:"Value"`
}

func (x *Environment) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var raw map[string]interface{}
	if err := unmarshal(&raw); err != nil {
		return err
	}
	if val, ok := raw["Type"]; ok {
		switch val {
		case "secret":
			x.Type = EnvironmentTypeSecret
		default:
			x.Type = EnvironmentTypeText
		}
	}
	if val, ok := raw["Name"]; ok {
		if v, ok := val.(string); ok {
			x.Name = v
		}
	}
	if val, ok := raw["Value"]; ok {
		if v, ok := val.(string); ok {
			x.Value = v
		}
	}
	return nil
}
