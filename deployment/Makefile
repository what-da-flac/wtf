SHELL = bash
$(shell touch .env .env.credentials)
include .env .env.credentials
export

docker-login:
	./scripts/docker-login.sh

docker-logs:
	docker compose logs -f gateway listener

docker-pull: docker-login
	docker compose pull

docker-start:
	docker compose up -d

docker-stop:
	docker compose down --remove-orphans

logs: docker-logs

pg:
	docker compose exec -it postgres psql

sql-assign-roles:
	docker compose exec postgres psql -f /scripts/assign-roles.sql

start:
	sudo service wtf start

stop:
	sudo service wtf stop

update: docker-pull stop start update-ui sql-assign-roles
	
update-ui:
	./scripts/update-ui.sh $(UI_VERSION)
