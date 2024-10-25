package i_am

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func ParseGroups(filename string) (ModelGroups, error) {
	log.Println("processing file:", filename)
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() { _ = file.Close() }()
	var b []*ModelGroup
	if err := yaml.NewDecoder(file).Decode(&b); err != nil {
		return nil, err
	}
	for _, v := range b {
		if err := v.Validate(); err != nil {
			return nil, fmt.Errorf("error validating group: %s %w", v.Name, err)
		}
	}
	return b, nil
}

func ParseUsers(filename string) ([]*ModelUser, error) {
	log.Println("processing file:", filename)
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() { _ = file.Close() }()
	var b []*ModelUser
	if err := yaml.NewDecoder(file).Decode(&b); err != nil {
		return nil, err
	}
	return b, nil
}
