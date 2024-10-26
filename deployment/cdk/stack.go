package main

import (
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type StackProps struct {
	Name        string `yaml:"Name"`
	BaseDir     string `yaml:"BaseDir"`
	Description string `yaml:"Description"`
}

// ReadStacks loops over files in dir without going deep, for YAML files and parses
// StackProps array from each one of them.
func ReadStacks(dir string) ([]*StackProps, error) {
	const yamlExt = ".yaml"
	var stackProps []*StackProps
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	readFileFn := func(filename string) ([]*StackProps, error) {
		file, err := os.Open(filename)
		if err != nil {
			return nil, err
		}
		defer func() { _ = file.Close() }()
		var sps []*StackProps
		if err := yaml.NewDecoder(file).Decode(&sps); err != nil {
			return nil, err
		}
		return sps, nil
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if !strings.HasSuffix(entry.Name(), yamlExt) {
			continue
		}
		props, err := readFileFn(filepath.Join(dir, entry.Name()))
		if err != nil {
			panic(err)
		}
		stackProps = append(stackProps, props...)
	}
	return stackProps, nil
}
