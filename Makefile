COMPOSE=docker compose --env-file ./deployment/docker.env  -f ./deployment/docker-compose.yaml -p avito-shop
PG_URL=postgres://user:password@localhost:5432/shop?sslmode=disable

.PHONY: up
up:
	${COMPOSE} up -d

.PHONY: down
down:
	${COMPOSE} down

.PHONY: up-infra
up-infra:
	${COMPOSE} up -d postgres migrations

.PHONY: .migration-up
migration-up:
	$(eval PG_URL?=$(PG_URL))
	goose -dir ./migrations postgres "$(PG_URL)" up

.PHONY: .migration-down
migration-down:
	$(eval PG_URL?=$(PG_URL))
	goose -dir ./migrations postgres "$(PG_URL)" down

.PHONY: .migration-status
migration-status:
	$(eval PG_URL?=$(PG_URL))
	goose -dir ./migrations postgres "$(PG_URL)" status

.PHONY: .migration-create-sql
migration-create-sql:
	goose -dir ./migrations create $(filter-out $@,$(MAKECMDGOALS)) sql
