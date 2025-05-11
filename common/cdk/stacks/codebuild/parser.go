package codebuild

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func Parse(filename string) (*Model, error) {
	log.Println("processing file:", filename)
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() { _ = file.Close() }()
	b := &Model{}
	if err = yaml.NewDecoder(file).Decode(b); err != nil {
		return nil, err
	}
	if err = b.Validate(); err != nil {
		return nil, err
	}
	return b, nil
}
