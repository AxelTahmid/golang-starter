## ⚠️ WIP ⚠️

Things yet to add

-   Docgen in .md for routes

## Get Started

Copy `.env.example` to `.env` to get started & provide appropiate values

```sh
cp .env.example .env
```

Generate TLS cert for serving https locally

```sh
make tls
```

Generate ECDSA public-private key pair for JWT Auth

```sh
make jwt
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
