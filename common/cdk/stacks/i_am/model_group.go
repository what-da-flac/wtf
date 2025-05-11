package i_am

import (
	"fmt"

	"github.com/what-da-flac/wtf/common/cdk/stacks/common"
)

type ModelGroup struct {
	ManagedPolicies []string         `yaml:"ManagedPolicies"`
	Name            string           `yaml:"Name"`
	Policies        []*common.Policy `yaml:"Policies"`
}

func (x *ModelGroup) Validate() error {
	if x.Name == "" {
		return fmt.Errorf("missing group name")
	}
	for _, policy := range x.Policies {
		if err := policy.Validate(); err != nil {
			return err
		}
	}
	return nil
}
