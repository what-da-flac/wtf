.PHONY: swagger-install swagger-gen swagger-clean swagger-bundle

GO_PACKAGE_NAME=golang
SRC_DIR_GO:=gen/$(GO_PACKAGE_NAME)
OPEN_API_VERSION=v2.4.1
OPEN_API_OUTPUT:=openapi-gen.yaml

swagger-bundle:
	redocly bundle open-api/openapi.yaml --output $(OPEN_API_OUTPUT)

swagger-clean:
	mkdir -p $(SRC_DIR_GO)
	rm -f $(SRC_DIR_GO)/types.go
	echo 'package go' > $(SRC_DIR_GO)/types.go

swagger-gen: swagger-clean swagger-bundle swagger-gen-go

swagger-gen-go:
	oapi-codegen --package $(GO_PACKAGE_NAME) --generate types,client,std-http $(OPEN_API_OUTPUT) > "$(SRC_DIR_GO)/types.go"

swagger-gen-node:
	openapi-generator-cli generate -i $(OPEN_API_OUTPUT) -g typescript-axios -o clients/node

swagger-gen-python:
	datamodel-codegen --input $(OPEN_API_OUTPUT) --input-file-type openapi --output models.py

swagger-install:
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@$(OPEN_API_VERSION)
	npm i -g @redocly/cli@latest
	npm install -g @openapitools/openapi-generator-cli
	pip install datamodel-code-generator
