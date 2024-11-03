.PHONY: ci _next-tag next-tag-cdk next-tag-docker next-tag-gateway next-tag-lambda next-tag-ui

ci:
	make -C services/ ci
	make -C go-common/ ci
	make -C openapi/ ci

_next-tag:
	git tag | go run ./gateway next-version $(ARGS)

next-tag-cdk:
	make _next-tag ARGS=cdk

next-tag-docker:
	make _next-tag ARGS=docker

next-tag-gateway:
	make _next-tag ARGS=gateway

next-tag-lambda:
	make _next-tag ARGS=lambda

next-tag-ui:
	make _next-tag ARGS=ui
