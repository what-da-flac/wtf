package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/what-da-flac/wtf/go-common/cdk/stacks"
)

// Deploy looks up all YAML files within baseDir parameter
// following a naming convention.
// Configuration will be applied as described on each YAML file.
func Deploy(stack awscdk.Stack, baseDir string) {
	var (
		codebuildFilenames []string
		ecrFilenames       []string
		groupFilenames     []string
		lambdaFilenames    []string
		s3Filenames        []string
		sqsFilenames       []string
		userFilenames      []string
	)
	listFilesFn := func(dirname string) []string {
		var res []string
		files, err := os.ReadDir(dirname)
		if err != nil {
			log.Println("error reading directory", dirname, err.Error())
			panic(err)
		}
		for _, file := range files {
			res = append(res, filepath.Join(dirname, file.Name()))
		}
		return res
	}
	dirs := listFilesFn(baseDir)
	for _, dir := range dirs {
		files := listFilesFn(dir)
		currDir := strings.TrimPrefix(dir, baseDir+"/")
		switch currDir {
		case "codebuild":
			codebuildFilenames = files
		case "ecr":
			ecrFilenames = files
		case "group":
			groupFilenames = files
		case "lambda":
			lambdaFilenames = files
		case "s3":
			s3Filenames = files
		case "sqs":
			sqsFilenames = files
		case "user":
			userFilenames = files
		}
	}
	stacks.RunAll(stack,
		codebuildFilenames, ecrFilenames,
		groupFilenames, lambdaFilenames,
		userFilenames, s3Filenames, sqsFilenames)
}
