# # Load all values from .env and export them, within makefile commands
# ifneq (,$(wildcard ./.env))
#     include .env
#     export
# endif

# self-seigned tls for local dev only
tls:
	cd ./cert && \
	openssl req -nodes -newkey rsa:2048 -new -x509 -keyout tls.key -out tls.crt -days 365 \
	-subj "//C=BD/ST=Dhaka/L=Dhaka/O=Golang/CN=localhost"

jwt:
	cd ./cert && \
	openssl ecparam -genkey -name prime256v1 -noout -out jwt-pvt.pem && \
	openssl ec -in jwt-pvt.pem -pubout -out jwt-pub.pem

deps:
	go mod download

deps-upgrade:
	go get -u -t -d -v ./...

deps-cleancache:
	go clean -modcache

tidy:
	go mod tidy
	go mod download

run: 
	tidy
	go run ./cmd/golang-starter/main.go

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

# Database migration scripts
db:
	docker compose --profile tools run --rm goose status
# Migrate the DB to the most recent version available
db-up:
	docker compose --profile tools run --rm goose up
# Roll back the version by 1
db-down:
	docker compose --profile tools run --rm goose down
# Re-run the latest migration
db-redo:
	docker compose --profile tools run --rm goose redo
# Roll back all migrations
db-reset:
	docker compose --profile tools run --rm goose reset
# Check migration files without running them
db-validate:
	docker compose --profile tools run --rm goose validate
# Creates new migration file with the current sequence 
# example: make migrate-create f=xxx
db-create:
	docker compose --profile tools run --rm goose create $(f) sql

lint:
	docker run -t --rm -v $(PWD):/app -w /app golangci/golangci-lint:v1.59-alpine golangci-lint run -v

lint-cache:
	docker run --rm -v $(PWD):/app -v ~/.cache/golangci-lint/v1.59-alpine:/root/.cache -w /app golangci/golangci-lint:v1.59.1 golangci-lint run -v