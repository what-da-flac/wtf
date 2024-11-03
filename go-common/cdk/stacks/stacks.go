package stacks

import (
	"log"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/what-da-flac/wtf/go-common/cdk/stacks/codebuild"
	"github.com/what-da-flac/wtf/go-common/cdk/stacks/ecr"
	"github.com/what-da-flac/wtf/go-common/cdk/stacks/i_am"
	"github.com/what-da-flac/wtf/go-common/cdk/stacks/lambda"
	"github.com/what-da-flac/wtf/go-common/cdk/stacks/s3"
	"github.com/what-da-flac/wtf/go-common/cdk/stacks/sqs"
)

func RunAll(
	stack awscdk.Stack,
	codebuildFilenames []string,
	ecrFilenames []string,
	groupFilenames []string,
	lambdaFilenames []string,
	userFilenames []string,
	s3Filenames []string,
	sqsFilenames []string,
) {
	RunCodebuild(stack, codebuildFilenames...)
	RunEcr(stack, ecrFilenames...)
	RunIamGroups(stack, groupFilenames, userFilenames)
	RunS3(stack, s3Filenames...)
	RunSQS(stack, sqsFilenames...)
	RunLambda(stack, lambdaFilenames...)
}

func RunCodebuild(stack awscdk.Stack, filenames ...string) {
	for _, filename := range filenames {
		build, err := codebuild.Parse(filename)
		if err != nil {
			panic(err)
		}
		codebuild.Build(stack, build)
	}
}

func RunEcr(stack awscdk.Stack, filenames ...string) {
	var builds []*ecr.Model
	for _, filename := range filenames {
		build, err := ecr.Parse(filename)
		if err != nil {
			panic(err)
		}
		builds = append(builds, build...)
	}
	for _, build := range builds {
		ecr.Build(stack, build)
	}
}

func RunIamGroups(stack awscdk.Stack, groupFilenames, userFilenames []string) {
	var (
		modelGroups []*i_am.ModelGroup
		modelUsers  []*i_am.ModelUser
	)
	for _, filename := range groupFilenames {
		groups, err := i_am.ParseGroups(filename)
		if err != nil {
			panic(err)
		}
		modelGroups = append(modelGroups, groups...)
	}
	for _, filename := range userFilenames {
		users, err := i_am.ParseUsers(filename)
		if err != nil {
			panic(err)
		}
		modelUsers = append(modelUsers, users...)
	}
	i_am.Build(stack, modelGroups, modelUsers)
}

func RunLambda(stack awscdk.Stack, filenames ...string) {
	var lambdas []*lambda.Model
	for _, filename := range filenames {
		items, err := lambda.Parse(filename)
		if err != nil {
			panic(err)
		}
		lambdas = append(lambdas, items...)
	}
	for _, l := range lambdas {
		lambda.Build(stack, l)
	}
}

func RunS3(stack awscdk.Stack, filenames ...string) {
	for _, filename := range filenames {
		builds, err := s3.Parse(filename)
		if err != nil {
			log.Println("error parsing file:", filename, err)
			panic(err)
		}
		for _, build := range builds {
			s3.Build(stack, build)
		}
	}
}

func RunSQS(stack awscdk.Stack, filenames ...string) {
	for _, filename := range filenames {
		build, err := sqs.Parse(filename)
		if err != nil {
			panic(err)
		}
		for _, v := range build {
			sqs.Build(stack, v)
		}
	}
}
