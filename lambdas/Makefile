.PHONY: ci docker-build-all docker-push-all docker-tag-all

SHELL := /bin/bash

LAMBDAS=torrent-metadata torrent-download
BASE_GO_BUILD_IMAGE=public.ecr.aws/docker/library/golang
BASE_GO_VERSION=1.23
GIT_TAG=$(shell git describe --tags --abbrev=0)

ci: test-all

docker-build-all:
	@$(foreach lambda, $(LAMBDAS), \
		docker build --progress plain \
		--build-arg BASE_GO_BUILD_IMAGE=$(BASE_GO_BUILD_IMAGE) \
		--build-arg BASE_GO_VERSION=$(BASE_GO_VERSION) \
		--build-arg GITHUB_ACCESS_TOKEN=$(GITHUB_ACCESS_TOKEN) \
		--build-arg LAMBDA_NAME=$(lambda) \
		-t $(lambda):$(GIT_TAG) \
		-f $(lambda)/Dockerfile \
		.. ; \
	)

docker-push-all:
	@$(foreach lambda, $(LAMBDAS), \
		docker push $(ACCOUNT_ID).dkr.ecr.us-east-2.amazonaws.com/$(lambda):$(GIT_TAG); \
	)

docker-tag-all:
	@$(foreach lambda, $(LAMBDAS), \
		docker tag $(lambda):$(GIT_TAG) $(ACCOUNT_ID).dkr.ecr.us-east-2.amazonaws.com/$(lambda):$(GIT_TAG); \
	)

test-all:
	go test -cover  ./...
