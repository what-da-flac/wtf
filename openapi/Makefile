.PHONY: ci
.PHONY: swagger-install swagger-gen swagger-clean swagger-bundle
.PHONY: test


SRC_DIR:=models
OPEN_API_VERSION=v2.2.0
OPEN_API_OUTPUT:=openapi-gen.yaml

build:
	go build

ci: test

clean:
	rm -rf $(GEN_DIR)/*
	mkdir -p $(GEN_DIR)

swagger-bundle:
	redocly bundle open-api/openapi.yaml --output $(OPEN_API_OUTPUT)

swagger-clean:
	mkdir -p $(SRC_DIR)
	rm -f $(SRC_DIR)/types.go
	echo 'package models' > $(SRC_DIR)/types.go

swagger-gen: swagger-clean swagger-bundle
	oapi-codegen --package models --generate types,client,std-http $(OPEN_API_OUTPUT) > $(SRC_DIR)/types.go

swagger-install:
	go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@$(OPEN_API_VERSION)
	npm i -g @redocly/cli@latest

test:
	go test -short -cover ./...
