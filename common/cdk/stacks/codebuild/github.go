package codebuild

import "fmt"

type WebhookFilterType int

const (
	WebhookPullRequest WebhookFilterType = iota
	WebhookTagRelease  WebhookFilterType = iota
)

type Github struct {
	CodebuildScriptPath string            `yaml:"CodebuildScriptPath"`
	Repo                string            `yaml:"Repo"`
	Owner               string            `yaml:"Owner"`
	Filter              WebhookFilterType `yaml:"-"`
	FilterString        string            `json:"Filter"`
	PatternMatching     string            `yaml:"PatternMatching"`
}

func (x *Github) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var raw map[string]interface{}
	if err := unmarshal(&raw); err != nil {
		return err
	}
	if val, ok := raw["Filter"]; ok {
		switch val {
		case "pull-request":
			x.Filter = WebhookPullRequest
		case "tag-release":
			x.Filter = WebhookTagRelease
		default:
			return fmt.Errorf("unknown filter: %v", val)
		}
	}
	if val, ok := raw["CodebuildScriptPath"]; ok {
		if v, ok := val.(string); ok {
			x.CodebuildScriptPath = v
		}
	}
	if val, ok := raw["Repo"]; ok {
		if v, ok := val.(string); ok {
			x.Repo = v
		}
	}
	if val, ok := raw["Owner"]; ok {
		if v, ok := val.(string); ok {
			x.Owner = v
		}
	}
	if val, ok := raw["PatternMatching"]; ok {
		if v, ok := val.(string); ok {
			x.PatternMatching = v
		}
	}
	return nil
}
