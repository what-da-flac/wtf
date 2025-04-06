.PHONY: install gen clean bundle

GEN_DIR=gen
GO_PACKAGE_NAME=golang
DIR_GO:=$(GEN_DIR)/$(GO_PACKAGE_NAME)
DIR_NODE:=$(GEN_DIR)/node
DIR_PYTHON:=$(GEN_DIR)/py
DIR_TS:=$(GEN_DIR)/ts
OPEN_API_VERSION=v2.4.1
OPEN_API_OUTPUT:=openapi-gen.yaml

bundle:
	redocly bundle open-api/openapi.yaml --output $(OPEN_API_OUTPUT)

clean-go:
	mkdir -p $(DIR_GO)
	rm -f $(DIR_GO)/types.go
	echo 'package go' > $(DIR_GO)/types.go

clean-node:
	rm -rf $(DIR_NODE)
	mkdir -p $(DIR_NODE)

clean-python:
	rm -rf $(DIR_PYTHON)
	mkdir -p $(DIR_PYTHON)

clean-ts:
	rm -rf $(DIR_TS)
	mkdir -p $(DIR_TS)

gen: bundle gen-go gen-node gen-python gen-ts

gen-go: clean-go
	oapi-codegen --package $(GO_PACKAGE_NAME) --generate types,client,std-http $(OPEN_API_OUTPUT) > "$(DIR_GO)/types.go"

gen-node: clean-node
	npx openapi-typescript-codegen --input $(OPEN_API_OUTPUT) --output $(DIR_NODE) --client axios --useOptions --useUnionTypes

gen-python: clean-python
	datamodel-codegen --input $(OPEN_API_OUTPUT) --input-file-type openapi --output $(DIR_PYTHON)/models.py

gen-ts: clean-ts
	npx @openapitools/openapi-generator-cli generate  -i $(OPEN_API_OUTPUT) -g typescript-axios -o $(DIR_TS)

install:
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@$(OPEN_API_VERSION)
	npm i -g @redocly/cli@latest
	npm i -g @openapitools/openapi-generator-cli
	npm i -g openapi-typescript-codegen@0.29.0 --force
	pip install datamodel-code-generator
