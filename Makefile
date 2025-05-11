.PHONY: ci _next-tag next-tag-cdk next-tag-docker next-tag-service next-tag-ui
.PHONY: local-build local-logs local-pg local-start local-stop
.PHONY: push-tag-service push-tag-cdk push-tag-docker push-tag-ui

ci: test-all
	@make -C ui/ install ci

local-build:
	docker compose build

local-start:
	docker compose up -d

local-stop:
	docker compose down --remove-orphans

local-logs:
	docker compose logs -f

test-all:
	go test -short -cover ./common/...
	go test -short -cover ./openapi/...
	@make -C services/ ci
