package ecr

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func Parse(filename string) ([]*Model, error) {
	var b []*Model
	log.Println("processing file:", filename)
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() { _ = file.Close() }()
	if err := yaml.NewDecoder(file).Decode(&b); err != nil {
		return nil, err
	}
	for _, m := range b {
		if err := m.Validate(); err != nil {
			return nil, err
		}
	}
	return b, nil
}
