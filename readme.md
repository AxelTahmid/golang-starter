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
-   `prometheus/promhttp` for default metrics
-   `Makefile` commands for toolchain

## Folder Structure

The project follows below folder structure:

```go
project-root/
    ├── api/                     # API-related code (e.g. REST)
    │   ├── middleware/          # Middleware for HTTP requests
    │   │   └── ...
    │   ├── routes.go            # All Application routes
    │   └── server.go            # Server Setup
    ├── app/                     # Encapsulted Applicaton Logic
    │   ├── auth/
    │   │   ├── handler.go       # Http handlers specific to auth
    │   │   ├── router.go        # Http routes specific to auth
    │   │   └── types.go         # Types routes specific to auth
    │   └── ...                  # Other application logic folders
    ├── cert/                    # Certificates & Keys
    │   └── ...
    ├── cmd/
    │   ├── main.go              # Server entrypoint
    │   └── ...                  # Other application entrypoints
    ├── config/                  # Configuration logic
    │   ├── config.go
    │   └── ...
    ├── db/
    │   ├── migrations/          # Migrations files for goose
    │   │   └── *.sql
    │   ├── db.go                # Database connection setup
    │   ├── logger.go            # Database logger adapter
    │   └── ...                  # Database model files
    ├── docs/                    # Project documentations
    │   ├── bruno/               # Bruno collection for exploring api
    │   │   └── ...
    │   └── ...
    ├── pkg/                     # Public, reusable packages
    │   ├── <name>/
    │   │   └── ...              # Public package code
    │   └── ...
    ├── .gitignore               # Gitignore file
    ├── Makefile                 # Runnable Scripts
    ├── go.mod                   # Go module file
    ├── go.sum                   # Go module dependencies file
    └── README.md                # Project README
```

## Building Binaries

use below command to build binaries targetted towards supported `os` & `architechture` . A github workflow file is also included

```sh
make build os=<OPERATING SYSTEM> arch=<ARCHITECHTURE>
```

## Get Started - Development

### Environment

Ensure following tools available in your machine

-   Docker >= v27
-   Docker Compose >=v2.29
-   OpenSSL >= 1.1.1

### Server

To run the development server use below command:

```sh
make init
```

This command will

-   Copy `.env.example` to `.env`, filling in default values
-   Generate TLS cert for serving https locally.
-   Generate ECDSA public-private key pair for JWT Auth locally.
-   Start development server in hot reload mode
-   Migrate database to latest migration

After running above command for the first time, you should only start development server in hot reload mode using below command

```sh
make dev
```

n.b. In production environment, mount your certificates in `cert` directory, ensure proper name & path supplied in env.

### Exploring APIs

[Bruno](https://github.com/usebruno/bruno) is an Opensource IDE For Exploring and Testing Api's, lightweight alternative to postman/insomnia. Open Directory `docs/bruno` from the bruno app & the collection will show up. Choose the environment `Local`.

Download & install bruno -> [click here](https://www.usebruno.com/downloads).

### Other Commands

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

Generate TLS cert for serving https locally.

```sh
make tls
```

Generate ECDSA public-private key pair for JWT Auth locally.

```sh
make jwt
```

Run linter

```sh
make lint
```
