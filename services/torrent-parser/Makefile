.PHONY: build ci clean docker-build docker-start docker-stop
.PHONY: golang-imports golang-imports-install install
.PHONY: lint lint-clean lint-install
.PHONY: migrate-install migrate-new mock-clean mock-gen mock-install
.PHONY: sql-assign-roles
.PHONY: test test-html-output

DOCKER_IMAGE_NAME=gateway
LDFLAGS=-ldflags "-X=main.Version=$(SERVICE_VERSION)"
LINT_VERSION=v1.61.0
MIGRATE_VERSION=v4.18.1
MOQ_VERSION=v0.5.0
SERVICE_NAME := "service"
SERVICE_VERSION := $(shell git describe --tags --abbrev=0)
SHELL := /bin/bash
GIT_TAG=$(shell git describe --tags --abbrev=0 --always)
SRC_DIR:=internal/models

build:
	CGO_ENABLED=0 go build -buildvcs=false -a --installsuffix cgo -o $(SERVICE_NAME) $(LDFLAGS) && chmod 777 $(SERVICE_NAME)

ci: lint mock-gen test

clean:
	rm -rf $(GEN_DIR)/*
	mkdir -p $(GEN_DIR)

docker-build:
	docker build \
	--build-arg SERVICE_VERSION=$(SERVICE_VERSION) \
	--build-arg GITHUB_ACCESS_TOKEN=$(GITHUB_ACCESS_TOKEN) \
	-t $(DOCKER_IMAGE_NAME):$(GIT_TAG) \
	-f Dockerfile \
	..

docker-start:
	docker compose up -d

docker-stop:
	docker compose down --remove-orphans

golang-imports:
	goimports -w .

golang-imports-install:
	go install golang.org/x/tools/cmd/goimports@latest

install:
	go install

lint: mock-clean
	golangci-lint run --config golangci.yaml --timeout 10m

lint-clean:
	golangci-lint cache clean

lint-install:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin $(LINT_VERSION)

migrate-new:
	migrate create -ext sql -dir internal/assets/files/migrations $(NAME)

migrate-install:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@$(MIGRATE_VERSION)

mock-clean:
	rm -rf mocks
	mkdir -p mocks

mock-gen: mock-clean
	rm -rf mocks
	mkdir -p mocks
	go generate ./internal/interfaces/...

mock-install:
	go install github.com/matryer/moq@$(MOQ_VERSION)

sql-assign-roles:
	docker compose exec postgres psql -f /scripts/assign-roles.sql

test:
	go test -short -cover ./...

test-html-output:
	go test -short -coverprofile=c.out ./... && go tool cover -html=c.out && rm -f c.out
