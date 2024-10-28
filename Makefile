.PHONY: ci

ci:
	make -C gateway/ ci
	make -C go-common/ ci
	make -C lambdas/ ci
	make -C openapi/ ci