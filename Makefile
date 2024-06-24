# Load all values from .env and export them
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

run: 
	go run ./cmd/main.go

tidy:
	go mod tidy
	go mod vendor

deps-upgrade:
	go get -u -t -d -v ./...

deps-cleancache:
	go clean -modcache

build:
	docker compose up -d --build --no-cache

up:
	docker compose -f docker-compose.dev.yml up -d

down:
	docker compose -f docker-compose.dev.yml down

dev:
	down up log

enter-db:
	docker exec -it db sh

log:
	docker logs --follow api

log-db:
	docker logs --follow db

migrate-build:
	docker build -t migrate --target migrate \
		--build-arg="DB_USER=${DB_USER}" \
		--build-arg="DB_PASSWORD=${DB_ROOT_PASSWORD}" \
		--build-arg="DB_NAME=${DB_NAME}" \
		.

migrate-up:
	migrate-build
	docker run --network appnetwork migrate 

# TODO
# migrate-down:
# 	migrate-build
# 	docker run --network appnetwork migrate down
