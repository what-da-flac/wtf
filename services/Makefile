.PHONY: ci build-all golang-imports login mock-all push-all tag-all test-all
.PHONY: local-build local-logs local-pg local-start local-stop local-pull pull

# AWS account
ACCOUNT_ID=160885250498
SHELL := /bin/bash
SERVICES=gateway magnet-parser torrent-info torrent-parser
BASE_GO_BUILD_IMAGE=public.ecr.aws/docker/library/golang
BASE_GO_VERSION=1.23
GIT_TAG=$(shell git describe --tags --abbrev=0 --always)

$(shell touch .env .env.credentials)
include .env .env.credentials
export

ci: test-all

build-%:
	docker build \
		--build-arg BASE_GO_BUILD_IMAGE=$(BASE_GO_BUILD_IMAGE) \
		--build-arg BASE_GO_VERSION=$(BASE_GO_VERSION) \
		--build-arg GITHUB_ACCESS_TOKEN=$(GITHUB_ACCESS_TOKEN) \
		--build-arg SERVICE_NAME=$* \
		-t $*:$(GIT_TAG) \
		-f $*/Dockerfile \
		..

build-all:
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

golang-imports:
	goimports -w .

local-build:
	docker compose build

local-start:
	docker compose up -d

local-stop:
	docker compose down --remove-orphans

local-pull: login
	docker compose pull

login:
	aws --profile wtf ecr get-login-password --region us-east-2 | docker login --username AWS --password-stdin $(ACCOUNT_ID).dkr.ecr.us-east-2.amazonaws.com

local-logs:
	docker compose logs -f $(SERVICES)

local-pg:
	docker compose exec -it postgres psql

mock-all:
	@$(foreach service, $(SERVICES), \
		cd $(service) && rm -rf mocks && mkdir -p mocks && go generate ./... && cd ..; \
	)

pull:
	docker compose pull

push-all:
	@$(foreach service, $(SERVICES), \
		docker push $(ACCOUNT_ID).dkr.ecr.us-east-2.amazonaws.com/$(service):$(GIT_TAG); \
	)

tag-all:
	@$(foreach service, $(SERVICES), \
		docker tag $(service):$(GIT_TAG) $(ACCOUNT_ID).dkr.ecr.us-east-2.amazonaws.com/$(service):$(GIT_TAG); \
	)

test-all: mock-all
	go test -cover  ./...
