## Golang Http Starter - ⚠️ WIP ⚠️

This is a monolithic http api starter with sane defaults. Features include:

-   `slog` for global logging
-   `envconfig` for app configuration parsing & validation
-   `air` for hot-reloading dev server
-   `chi/v5` for routing & middleware
-   `pgx/v5` for database connectivity with pooling
-   `migrate/migrate` for migrations without adding to app dependency
-   `bcrypt` fir password hashing
-   `golang-jwt/jwt/v5` for JWT Authentication
-   JSON request validation with `validator/v10` & response serialization
-   `Makefile` scripts for common usages

## Folder Structure

The project follows below folder structure:

```go
project-root/
    ├── cmd/
    │   ├── app-name/
    │   │   ├── main.go          # Application entry point
    │   │   └── ...              # Other application-specific files
    ├── cert/
    │   ├── ...                  # Certificates & Keys
    ├── config/
    │   ├── config.go            # Configuration logic
    │   └── ...
    ├── db/
    │   ├── migrations/
    │   │   ├── *.sql            # Migrations files
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
    │   ├── mypackage/
    │   │   ├── mypackage.go     # Public package code
    │   │   └── ...
    │   └── ...
    ├── domain/                  # Encapsulted Applicaton Logic
    │   ├── mypackage/
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

Ensure `Docker` & `Docker-Compose` is installed locally with compose v2 being available.

```sh
# v1 syntax
docker-compose ...
# v2 syntax
docker compose ...
```

Copy `.env.example` to `.env` to get started & provide appropiate values

```sh
cp .env.example .env
```

Generate TLS cert for serving https locally. In production environment, mount your certificate in `cert` directory with proper name.

```sh
make tls
```

Generate ECDSA public-private key pair for JWT Auth locally. In production environment, mount your keys in `cert` directory with proper name.

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
make db-up
```

revert migrations by 1 level

```sh
make db-down
```

create migration files

```sh
make db-create f=<filename>
```

## Building Binaries

use below command to build binaries targetted towards supported `os` & `architechture`

```sh
os=<OPERATING SYSTEM> arch=<ARCHITECHTURE> make build-release
```
