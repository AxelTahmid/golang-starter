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
migrate-down:
	docker run migrate/migrate -path /migrations -database "mysql://${DB_USER}:${DB_ROOT_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" down
