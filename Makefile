.PHONY: ci _next-tag next-tag-cdk next-tag-docker next-tag-service -tag-ui

ci:
	make -C services/ ci
	make -C go-common/ ci
	make -C openapi/ ci
	make -C ui/ ci

_next-tag:
	git tag | go run ./services/gateway next-version $(ARGS)

next-tag-cdk:
	make _next-tag ARGS=cdk

next-tag-docker:
	make _next-tag ARGS=docker

next-tag-service:
	make _next-tag ARGS=service

next-tag-ui:
	make _next-tag ARGS=ui
