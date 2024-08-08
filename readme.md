## Golang Api Starter

This is a monolithic http api starter with sane defaults. Features include:

-   `air` - hot-reloading dev server
-   `slog` - global logging
-   `chi/v5` - routing & middleware
-   `pgx/v5` - database connectivity with pooling
-   `goose` - database migrations without adding to app dependency
-   `bcrypt` - password hashing
-   `golang-jwt/jwt/v5` - jwt Authentication
-   `validator/v10` - incoming request payload validation
-   `envconfig` - app configuration parsing & validation
-   `Makefile` commands for tooling

## Folder Structure

The project follows below folder structure:

```go
project-root/
    ├── cmd/
    │   ├── <app-name>/
    │   │   ├── main.go          # Application entry point
    │   │   └── ...              # Other application-specific files
    ├── cert/
    │   ├── ...                  # Certificates & Keys
    ├── config/
    │   ├── config.go            # Configuration logic
    │   └── ...
    ├── db/
    │   ├── migrations/
    │   │   ├── *.sql            # Migrations files for goose
    │   │   ├── ...
    │   ├── db.go                # Database setup and access
    │   ├── logger.go            # Database logger adapter
    │   └── ...
    ├── api/                     # API-related code (e.g., REST)
    │   ├── middleware/
    │   │   ├── middleware.go    # Middleware for HTTP requests
    │   │   └── ...
    │   ├── routes.go            # All Application routes
    │   └── ...
    ├── pkg/                     # Public, reusable packages
    │   ├── <name>/
    │   │   ├── mypackage.go     # Public package code
    │   │   └── ...
    │   └── ...
    ├── domain/                  # Encapsulted Applicaton Logic
    │   ├── <nam>/
    │   │   ├── mypackage.go
    │   │   └── ...
    │   └── ...
    ├── docs/                    # Project documentation ( WIP )
    ├── .gitignore               # Gitignore file
    ├── Makefile                 # Runnable Scripts
    ├── go.mod                   # Go module file
    ├── go.sum                   # Go module dependencies file
    └── README.md                # Project README
```

## Get Started

Ensure `Docker` & `Docker-Compose` is installed locally with compose v2 available. Copy `.env.example` to `.env` to get started & fill in your own values

```sh
cp .env.example .env
```

In production environment, mount your certificates in `cert` directory, ensure proper name & path supplied in env. But for development, use below commands:

Generate TLS cert for serving https locally.

```sh
make tls
```

Generate ECDSA public-private key pair for JWT Auth locally.

```sh
make jwt
```

install & clean golang dependencies

```sh
make tidy
make deps
```

run the development server with hot reload

```sh
make dev
```

run migration of postgresql database, ensure `.env` values are correctly given

```sh
make migrate-up
```

revert migrations by 1 level

```sh
make migrate-down
```

create migration files

```sh
make migrate-create f=<filename>
```

## Building Binaries

use below command to build binaries targetted towards supported `os` & `architechture` . A github workflow file is also included

```sh
make build os=<OPERATING SYSTEM> arch=<ARCHITECHTURE>
```
