include .env
export

export PROJECT_ROOT = $(shell pwd)

env-up:
	@docker compose up -d postgres-app

env-down:
	@docker compose down postgres-app

env-cleanup:
	@read -p "Are you sure you want to delete the database? (y/n): " confirm; \
	if [ "$$confirm" = "y" ]; then \
		docker compose down postgres-app;  \
		rm -rf out/pgdata/;  \
		echo "Database deleted successfully."; \
	else \
		echo "Database deletion canceled."; \
	fi; \


env-port-forward:
	@docker compose up -d port-forwarder

env-port-close:
	@docker compose down port-forwarder


migration-create:

	@if [-z "$(seq)"]; then \
		echo "Error: Migration sequence number is required."; \
		exit 1; \
	fi;

	docker compose run --rm migrate-app \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migration-up:
	make migration-action action=up

migration-down:
	make migration-action action=down

migration-action:
	@if [-z "$(action)"]; then \
		echo "Error: Migration action is required (up or down)."; \
		exit 1; \
	fi;\

	docker compose run --rm migrate-app \
		-path /migrations \
		-database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@postgres-app:5432/$(POSTGRES_DB)?sslmode=disable" \
		"$(action)"
