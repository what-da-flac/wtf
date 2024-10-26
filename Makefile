.PHONY: ci

ci:
	MAKE -C go-common/ ci
	MAKE -C openapi/ ci