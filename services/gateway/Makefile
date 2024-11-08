.PHONY: migrate-new migrate-install

MIGRATE_VERSION=v4.18.1
SHELL := /bin/bash
SRC_DIR:=internal/models

migrate-new:
	migrate create -ext sql -dir internal/assets/files/migrations $(NAME)

migrate-install:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@$(MIGRATE_VERSION)

sql-assign-roles:
	docker compose exec postgres psql -f /scripts/assign-roles.sql
