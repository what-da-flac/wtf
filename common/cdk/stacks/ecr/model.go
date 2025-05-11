package ecr

import "fmt"

type Model struct {
	EmptyOnDelete   bool   `yaml:"EmptyOnDelete"`
	Mutable         bool   `yaml:"Mutable"`
	Name            string `yaml:"Name"`
	RemoveOnDestroy bool   `yaml:"RemoveOnDestroy"`
	UseDefaults     bool   `yaml:"UseDefaults"`
}

func (x *Model) Validate() error {
	if x.Name == "" {
		return fmt.Errorf("ecr name is required")
	}
	if x.UseDefaults {
		x.EmptyOnDelete = true
		x.Mutable = true
		x.RemoveOnDestroy = true
	}
	return nil
}
