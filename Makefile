.PHONY: ci

ci:
	make -C gateway/ ci
	make -C go-common/ ci
	make -C openapi/ ci