.PHONY: ci build-all golang-imports login mock-all push-all tag-all test-all
.PHONY: lint-install pull mock-all mock-install
.PHONY: sql-assign-roles tools-install

# AWS account
ACCOUNT_ID=160885250498
SHELL := /bin/bash
SERVICES=gateway media-process
NODE_SERVICES=
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

ci: test-all-go test-all-node

build-%:
	docker build --progress=plain \
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

test-all-go: mock-all
	go test -short -cover ./...

test-all-node:
	@$(foreach service, $(NODE_SERVICES), \
		$(MAKE) -C $(service) ci; \
	)

test-%: mock-%
	go test -short -cover ./$*/...

tools-install: lint-install mock-install