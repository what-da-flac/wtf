package i_am

import "fmt"

type ModelUser struct {
	Username string   `yaml:"Username"`
	Groups   []string `yaml:"Groups"`
}

func (x *ModelUser) Validate() error {
	if x.Username == "" {
		return fmt.Errorf("missing username")
	}
	return nil
}
