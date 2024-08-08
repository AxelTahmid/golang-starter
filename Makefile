.PHONY: build

# # Load all values from .env and export them, within makefile commands
# ifneq (,$(wildcard ./.env))
#     include .env
#     export
# endif

deps:
	go mod download
	go mod verify

deps-upgrade: 
	go get -u -t -d -v ./...

deps-cleancache: 
	go clean -modcache

tidy:
	go mod tidy

run: 
	go run ./cmd/golang-starter/main.go

build: 
	@if [ -z "$(os)" ] || [ -z "$(arch)" ]; then \
		echo "Error: Both 'os' and 'arch' variables must be set. Please use 'make build os=<value> arch=<value>'"; \
		exit 1; \
	fi
	CGO_ENABLED=0 GOOS=$(os) GOARCH=$(arch) go build -o ./bin/main ./cmd/golang-starter/main.go

up:
	@if [ ! -f .env ]; then \
        cp .env.example .env; \
    fi
	docker compose up -d

down:
	docker compose down

fresh:
	@if [ ! -f .env ]; then \
        cp .env.example .env; \
    fi
	docker compose down --remove-orphans
	docker compose build --no-cache
	docker compose up -d --build -V

dev: tidy down up log

exec-db: 
	docker exec -it db sh

log:
	docker logs -f api

log-db: 
	docker logs -f db

## Database migration scripts
db: 
	docker compose --profile tools run --rm goose status

# Migrate the DB to the most recent version available
migrate-up: 
	docker compose --profile tools run --rm goose up

# Roll back the version by 1
migrate-down: 
	docker compose --profile tools run --rm goose down

# Re-run the latest migration
migrate-redo: 
	docker compose --profile tools run --rm goose redo

# Roll back all migrations
migrate-fresh: 
	docker compose --profile tools run --rm goose reset

# Check migration files without running them
migrate-validate: 
	docker compose --profile tools run --rm goose validate

# Creates new migration file with the current sequence 
migrate-create:
	@if [ -z "$(filename)" ]; then \
		echo "Error: 'filename' variable must be set. Please use 'make migrate-create filename=<value>'"; \
		exit 1; \
	fi
	docker compose --profile tools run --rm goose create $(filename) sql

# self-seigned tls for local dev only
tls:
	cd ./cert && \
	openssl req -nodes -newkey rsa:2048 -new -x509 -keyout tls.key -out tls.crt -days 365 \
	-subj "//C=BD/ST=Dhaka/L=Dhaka/O=Golang/CN=localhost"

jwt:
	cd ./cert && \
	openssl ecparam -genkey -name prime256v1 -noout -out jwt-pvt.pem && \
	openssl ec -in jwt-pvt.pem -pubout -out jwt-pub.pem

lint: 
	docker run -t --rm -v $(PWD):/app -w /app golangci/golangci-lint:v1.59-alpine golangci-lint run -v

# WIP
# lint-cache: docker run --rm -v $(PWD):/app -v ~/.cache/golangci-lint/v1.59-alpine:/root/.cache -w /app golangci/golangci-lint:v1.59.1 golangci-lint run -v