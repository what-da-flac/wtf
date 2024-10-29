.PHONY: build ci clean
.PHONY: docker-logs docker-start docker-stop
.PHONY: golang-imports golang-imports-install
.PHONY: lint lint-install
.PHONY: protos
.PHONY: test test-html-output test-no-cache

GEN_DIR := "protos/gen"
LINT_VERSION=v1.61.0
SHELL := /bin/sh

$(shell touch .env)
include .env
export


build:
	CGO_ENABLED=0 go build -buildvcs=false

ci: lint test

clean:
	rm -rf $(GEN_DIR)/*
	mkdir -p $(GEN_DIR)

docker-logs:
	docker compose logs -f

docker-start:
	docker compose up -d

docker-stop:
	docker compose down --remove-orphans

golang-imports:
	goimports -w .

golang-imports-install:
	go install golang.org/x/tools/cmd/goimports@latest

lint:
	golangci-lint run --config golangci.yaml --timeout 10m

lint-install:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin $(LINT_VERSION)
test:
	go test -cover -short ./...

test-html-output:
	go test -short -coverprofile=c.out ./... && go tool cover -html=c.out && rm -f c.out

test-no-cache:
	go test -cover -count 1 ./...
