FROM golang:1.22-alpine as base

## Dev Runner
FROM base as dev

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

EXPOSE 3000

CMD ["air", "-c", ".air.toml"]

## Production Builder
FROM base as build

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 go build -o /app/bin/main /app/cmd/main.go

## Production Runner
FROM gcr.io/distroless/static-debian12 as prod

COPY --from=build /go/bin/main /
CMD ["/main"]