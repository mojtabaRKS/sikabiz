help:
	@echo "Available commands:"
	@echo ""
	@echo "Docker Services:"
	@echo "  make infra             - Start core infrastructure (postgres)"
	@echo "  make up                - Start all docker compose services in detached mode"
	@echo "  make build-up          - Build and start all services"
	@echo "  make build-no-cache    - Build services without using cache"
	@echo "  make down              - Stop and remove containers (use RUN_ARGS for specific services)"
	@echo "  make purge             - Stop containers and remove volumes (use RUN_ARGS for specific services)"
	@echo ""
	@echo "Monitoring:"
	@echo "  make ps                - Show running docker compose services"
	@echo "  make status            - Show service status (use RUN_ARGS for specific services)"
	@echo "  make log               - Follow logs (use RUN_ARGS for specific services, e.g. RUN_ARGS=app)"
	@echo ""
	@echo "Database Migration:"
	@echo "  make migrator-up       - Start migrator service"
	@echo "  make migrator-down     - Stop migrator service"
	@echo "  make migrate           - Run database migrations (up)"
	@echo "  make migrate-up        - Run migrations up"
	@echo "  make migrate-down      - Run migrations down"
	@echo ""
	@echo "Examples:"
	@echo "  make log RUN_ARGS=app            - Follow logs for app service"
	@echo "  make shell RUN_ARGS=app          - Open shell in app service"
	@echo "  make down RUN_ARGS=tester        - Stop only tester service"

env:
	@[ -e ./.env ] || cp -v ./.env.example ./.env

infra:
	docker compose up -d postgres

up:
	docker compose up -d

ps:
	docker compose ps

build-up:
	docker compose up --build -d

build-no-cache:
	docker compose build --no-cache

status:
	docker compose ps $(RUN_ARGS)

down:
	docker compose down --remove-orphans $(RUN_ARGS)

purge:
	docker compose down --remove-orphans --volumes $(RUN_ARGS)

log:
	docker compose logs -f $(RUN_ARGS)

generate-swagger:
	swag init -q --parseDependency --parseInternal -g router.go -d internal/api

lint:
	golangci-lint run

migrator-up:
	docker compose --profile tools up -d migrator

migrator-down:
	docker compose --profile tools down migrator

migrate-up:
	docker compose exec migrator ./sikabiz migrate up

migrate-down:
	docker compose exec migrator ./sikabiz migrate down

migrate:
	$(MAKE) migrate-up

import:
	docker compose exec server ./sikabiz import
