.PHONY: ci docker-build-all docker-push-all docker-tag-all
.PHONY: mock-all test-all

SHELL := /bin/bash

SERVICES=gateway
BASE_GO_BUILD_IMAGE=public.ecr.aws/docker/library/golang
BASE_GO_VERSION=1.23
GIT_TAG=$(shell git describe --tags --abbrev=0 --always)

ci: test-all

docker-build-all:
	@$(foreach service, $(SERVICES), \
		docker build --progress plain \
		--build-arg BASE_GO_BUILD_IMAGE=$(BASE_GO_BUILD_IMAGE) \
		--build-arg BASE_GO_VERSION=$(BASE_GO_VERSION) \
		--build-arg GITHUB_ACCESS_TOKEN=$(GITHUB_ACCESS_TOKEN) \
		--build-arg SERVICE_NAME=$(service) \
		-t $(service):$(GIT_TAG) \
		-f $(service)/Dockerfile \
		.. ; \
	)

docker-push-all:
	@$(foreach service, $(SERVICES), \
		docker push $(ACCOUNT_ID).dkr.ecr.us-east-2.amazonaws.com/$(service):$(GIT_TAG); \
	)

docker-tag-all:
	@$(foreach service, $(SERVICES), \
		docker tag $(service):$(GIT_TAG) $(ACCOUNT_ID).dkr.ecr.us-east-2.amazonaws.com/$(service):$(GIT_TAG); \
	)

mock-all:
	@$(foreach service, $(SERVICES), \
		cd $(service) && rm -rf mocks && mkdir -p mocks && go generate ./... && cd ..; \
	)

test-all: mock-all
	go test -cover  ./...
