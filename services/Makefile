.PHONY: ci build-all golang-imports login mock-all push-all tag-all test-all
.PHONY: local-build local-logs local-pg local-start local-stop local-pull
.PHONY: lint-install pull mock-all mock-install
.PHONY: sql-assign-roles tools-install

# AWS account
ACCOUNT_ID=160885250498
SHELL := /bin/bash
SERVICES=gateway
BASE_GO_BUILD_IMAGE=golang
BASE_GO_VERSION=1.24.1
GIT_TAG := $(or $(GIT_TAG), local-dev.0.0.1)
GO_ROOT_PACKAGE=github.com/what-da-flac/wtf/services
LINT_VERSION=v2.0.2
MOCK_VERSION=v0.5.3

$(shell touch .env .env.credentials)
include .env .env.credentials
export

ci-%: mock-% test-%
	go test -cover  ./$*/...

build-%:
	docker build --progress=plain --no-cache \
		--build-arg BASE_GO_BUILD_IMAGE=$(BASE_GO_BUILD_IMAGE) \
		--build-arg BASE_GO_VERSION=$(BASE_GO_VERSION) \
		--build-arg SERVICE_NAME=$* \
		--build-arg GIT_TAG=$(GIT_TAG) \
		--build-arg GO_ROOT_PACKAGE=$(GO_ROOT_PACKAGE) \
		-t $*:$(GIT_TAG) \
		-f $*/Dockerfile \
		..

build-all:
	@$(foreach service, $(SERVICES), \
		$(MAKE) build-$(service) ; \
	)

golang-imports:
	goimports -w .

lint:
	golangci-lint run --config ../golangci.yaml --timeout 10m

lint-install:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin $(LINT_VERSION)

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
	find . -iname "mocks" -exec rm -rf {} \; || true
	go generate ./...

mock-%:
	cd $* && rm -rf mocks && mkdir -p mocks && go generate ./...

mock-install:
	go install github.com/matryer/moq@$(MOCK_VERSION)

pull:
	docker compose pull

push-all:
	@$(foreach service, $(SERVICES), \
		docker push $(ACCOUNT_ID).dkr.ecr.us-east-2.amazonaws.com/$(service):$(GIT_TAG); \
	)

sql-assign-roles:
	docker compose exec postgres psql -f /scripts/assign-roles.sql

tag-all:
	@$(foreach service, $(SERVICES), \
		docker tag $(service):$(GIT_TAG) $(ACCOUNT_ID).dkr.ecr.us-east-2.amazonaws.com/$(service):$(GIT_TAG); \
	)

test-all: mock-all
	go test -short -cover ./...

test-%: mock-%
	go test -short -cover ./$*/...

tools-install: lint-install mock-install