.PHONY: ci _next-tag next-tag-cdk next-tag-docker next-tag-service next-tag-ui
.PHONY: push-tag-service push-tag-cdk push-tag-docker push-tag-ui

ci: test-all
	@make -C ui/ install ci

_next-tag:
	@echo $(shell git tag | go run ./services/gateway next-version $(ARGS))

next-tag-cdk:
	@$(MAKE) _next-tag ARGS=cdk

next-tag-docker:
	@$(MAKE) _next-tag ARGS=docker

next-tag-service:
	@$(MAKE) _next-tag ARGS=service

next-tag-ui:
	@$(MAKE) _next-tag ARGS=ui

push-tag-cdk:
	@TAG_NAME=$(shell $(MAKE) next-tag-cdk) && git tag $$TAG_NAME && git push --tags

push-tag-docker:
	@TAG_NAME=$(shell $(MAKE) next-tag-docker) && git tag $$TAG_NAME && git push --tags

push-tag-service:
	@TAG_NAME=$(shell $(MAKE) next-tag-service) && git tag $$TAG_NAME && git push --tags

push-tag-ui:
	@TAG_NAME=$(shell $(MAKE) next-tag-ui) && git tag $$TAG_NAME && git push --tags

test-all:
	go test -short -cover ./openapi/...
	go test -short -cover ./go-common/...
	@make -C services/ ci
