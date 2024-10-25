package common

import (
	"fmt"
	"strings"
)

type DockerRegistry int

const (
	DockerRegistryCustom DockerRegistry = iota
	DockerRegistryAWS    DockerRegistry = iota
)

type Docker struct {
	RegistryType DockerRegistry `yaml:"-"`
	Type         string         `yaml:"Type"`
	Url          string         `yaml:"Url"`
}

func (x *Docker) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var raw map[string]interface{}
	if err := unmarshal(&raw); err != nil {
		return err
	}
	if val, ok := raw["Type"]; ok {
		switch val {
		case "custom":
			x.RegistryType = DockerRegistryCustom
		case "aws":
			x.RegistryType = DockerRegistryAWS
		default:
			return fmt.Errorf("unknown type: %v", val)
		}
	}
	if val, ok := raw["Url"]; ok {
		if v, ok := val.(string); ok {
			x.Url = v
		}
	}
	return nil
}

func (x *Docker) RepositoryTagName() (name, tag string) {
	const sep = ":"
	if x.RegistryType != DockerRegistryCustom {
		return "", ""
	}
	values := strings.Split(x.Url, sep)
	if len(values) != 2 {
		return
	}
	return values[0], values[1]
}
