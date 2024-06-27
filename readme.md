## TODOs

-   Database connectivity - pgxpool
-   Request Validation - validator/v10
-   Docgen in .md for routes

## Get Started

Generate TLS cert for serving

```sh
make tls
```

install & clean dependencies

```sh
make tidy
```

run the development server with hot realod

```sh
make dev
```

-   Make migration : `make migrate-create filename=xxx`
