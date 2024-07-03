# Load all values from .env and export them
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

# self-seigned tls for local dev only
tls:
	cd ./cert && \
	openssl req -nodes -newkey rsa:2048 -new -x509 -keyout tls.key -out tls.crt -days 365 \
	-subj "//C=BD/ST=Dhaka/L=Dhaka/O=Golang/CN=localhost"

jwt:
	cd ./cert && \
	openssl ecparam -genkey -name prime256v1 -noout -out private.pem && \
	openssl ec -in private.pem -pubout -out public.pem

deps:
	go mod download

deps-upgrade:
	go get -u -t -d -v ./...

deps-cleancache:
	go clean -modcache

run: 
	go mod download
	go run ./cmd/main.go

tidy:
	go mod tidy
	go mod download

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

# https://github.com/golang-migrate/migrate/issues/282#issuecomment-530743258
# make migrate-force v=xxx
migrate-force:
	docker compose --profile tools run --rm migrate force ${v}

# make migrate-create f=xxx
migrate-create:
	docker compose --profile tools run --rm migrate create -ext sql -dir /migrations $(f)

lint:
	docker run -t --rm -v $(PWD):/app -w /app golangci/golangci-lint:v1.59-alpine golangci-lint run -v

lint-cache:
	docker run --rm -v $(PWD):/app -v ~/.cache/golangci-lint/v1.59-alpine:/root/.cache -w /app golangci/golangci-lint:v1.59.1 golangci-lint run -v