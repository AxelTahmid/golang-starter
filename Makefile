# Load all values from .env and export them
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

# tls for local dev only, in deployment certbot is used. 
tls:
	cd ./cert && \
	openssl req -nodes -newkey rsa:2048 -new -x509 -keyout tls.key -out tls.crt -days 365 \
	-subj "//C=BD/ST=Dhaka/L=Dhaka/O=Golang/CN=localhost"

run: 
	deps
	go run ./cmd/main.go

tidy:
	go mod tidy
	deps

deps:
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
	docker compose --profile tools run --rm migrate up

migrate-down:
	docker compose --profile tools run --rm migrate down 1

# make migrate-create filename=xxx
migrate-create:
	docker compose --profile tools run --rm migrate create -ext sql -dir /migrations $(filename)