include .env
export

export PROJECT_ROOT=${shell pwd}

SHELL := /bin/bash

.PHONY: env-up env-down env-port-forward env-port-close env-cleanup \
        migrate-create migrate-up migrate-down migrate-action todoapp-run

env-up:
	docker compose up -d todoapp-postgres

env-down:
	@docker compose rm -sf todoapp-postgres

env-port-forward: ## env: Открыть порты сервисов окружения
	@docker compose up -d port-forwarder

env-port-close: ## env: Закрыть порты сервисов окружения
	@docker compose rm -sf port-forwarder

env-cleanup:
	@read -p "Очистить все volume файлы окружения? Опасность утери данных. [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
	  docker compose rm -sf todoapp-postgres && \
	  rm -rf ${PROJECT_ROOT}/out/pgdata && \
	  echo "Файлы окружения очищены"; \
	else \
	  echo "Очистка окружения отменена"; \
	fi

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Отсутствует необходимый параметр seq. Пример: make migrate-create seq=init"; \
		exit 1; \
	fi; \
	docker compose run --rm todoapp-postgresql-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migrate-up:
	@$(MAKE) migrate-action action=up

migrate-down:
	@$(MAKE) migrate-action action="down 1"

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "Отсутствует необходимый параметр action. Пример: make migrate-action action=up"; \
		exit 1; \
	fi; \
	docker compose run --rm todoapp-postgresql-migrate \
		-path /migrations \
		-database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@todoapp-postgres:5432/$(POSTGRES_DB)?sslmode=disable" \
		$(action)


todoapp-run:
	@export LOGGER_FOLDER=${PROJECT_ROOT}/out/logs && \
	export POSTGRES_HOST=localhost && \
	go mod tidy && \
	go run  ${PROJECT_ROOT}/cmd/todoapp/main.go
