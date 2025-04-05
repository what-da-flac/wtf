.PHONY: ci
.PHONY: swagger-install swagger-gen swagger-clean swagger-bundle
.PHONY: test


SRC_DIR:=models
OPEN_API_VERSION=v2.4.1
OPEN_API_OUTPUT:=openapi-gen.yaml

build:
	go build

ci: lint test

clean:
	rm -rf $(GEN_DIR)/*
	mkdir -p $(GEN_DIR)

lint:
	golangci-lint run --config golangci.yaml --timeout 10m

swagger-bundle:
	redocly bundle open-api/openapi.yaml --output $(OPEN_API_OUTPUT)

swagger-clean:
	mkdir -p $(SRC_DIR)
	rm -f $(SRC_DIR)/types.go
	echo 'package models' > $(SRC_DIR)/types.go

swagger-gen: swagger-clean swagger-bundle swagger-gen-go swagger-gen-node swagger-gen-python

swagger-gen-go:
	oapi-codegen --package models --generate types,client,std-http $(OPEN_API_OUTPUT) > $(SRC_DIR)/types.go

swagger-gen-node:
	openapi-generator-cli generate -i $(OPEN_API_OUTPUT) -g typescript-axios -o clients/node

swagger-gen-python:
	datamodel-codegen --input $(OPEN_API_OUTPUT) --input-file-type openapi --output models.py

swagger-install:
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@$(OPEN_API_VERSION)
	npm i -g @redocly/cli@latest
	npm install -g @openapitools/openapi-generator-cli
	pip install datamodel-code-generator

test:
	go test -short -cover ./...
