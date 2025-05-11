package common

import (
	"fmt"
	"strings"

	"github.com/what-da-flac/wtf/common/identifiers"
)

type Policy struct {
	Action    string   `yaml:"Action"`
	Name      string   `yaml:"Name"`
	Resources []string `yaml:"Resources"`
}

func (x *Policy) Validate() error {
	if x.Action == "" {
		return fmt.Errorf("missing policy action")
	}
	if x.Name == "" {
		return fmt.Errorf("missing policy name")
	}
	if len(x.Resources) == 0 {
		return fmt.Errorf("no resources defined for policy: %s", x.Name)
	}
	return nil
}

// GenRandomName creates a pseudo random name with a fixed prefix.
func GenRandomName(prefix string) string {
	const (
		offset = 10
		sep    = "-"
	)
	rndStr := identifiers.NewIdentifier().UUIDv4()[:offset]
	return prefix + "_" + strings.ReplaceAll(rndStr, sep, "")
}
