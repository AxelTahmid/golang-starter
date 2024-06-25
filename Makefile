# Load all values from .env and export them
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

run: 
	go run ./cmd/main.go

tidy:
	go mod tidy
	go mod download

deps-upgrade:
	go get -u -t -d -v ./...

deps-cleancache:
	go clean -modcache

build:
	docker compose up -d --build --no-cache

up:
	docker compose up -d

down:
	docker compose down

dev: down up log

enter-db:
	docker exec -it db sh

log:
	docker logs -f api

log-db:
	docker logs -f db

migrate-up:
	docker compose --profile tools run --rm -v migrate up

migrate-down:
	docker compose --profile tools run --rm -v migrate down 1

migrate-create:
	docker compose --profile tools run --rm -v migrate create -ext sql -dir /migrations $(filename)