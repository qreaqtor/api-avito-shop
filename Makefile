COMPOSE=docker compose --env-file ./deployment/docker.env  -f ./deployment/docker-compose.yaml -p avito-shop

.PHONY: up
up:
	${COMPOSE} up -d

.PHONY: down
down:
	${COMPOSE} down

.PHONY: up-infra
up-infra-dev:
	${COMPOSE} up -d postgres
