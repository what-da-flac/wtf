package s3

import (
	"github.com/what-da-flac/wtf/common/cdk/stacks/common"
)

// Model defines operations to configure S3 bucket.
type Model struct {
	AutoDeleteObjects    bool                     `yaml:"AutoDeleteObjects"`
	BlockPublicAccess    bool                     `yaml:"BlockPublicAccess"`
	EnforceSSL           bool                     `yaml:"EnforceSSL"`
	ExpirationDays       int                      `yaml:"ExpirationDays"`
	InlinePolicies       map[string]common.Policy `yaml:"InlinePolicies"`
	ManagedPolicies      []string                 `yaml:"ManagedPolicies"`
	Name                 string                   `yaml:"Name"`
	RemoveOnDestroy      bool                     `yaml:"RemoveOnDestroy"`
	Versioned            bool                     `yaml:"Versioned"`
	WebsiteIndexDocument string                   `yaml:"WebsiteIndexDocument"`
}
