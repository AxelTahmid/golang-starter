FROM golang:1.22-alpine as dev

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 3000

CMD ["air", "-c", ".air.toml"]


FROM migrate/migrate as migrate

COPY ./database/migrations /migrations

ARG DB_NAME
ARG DB_USER
ARG DB_PASSWORD

ENTRYPOINT [ "migrate", "-path", "/migrations", "-database"]

CMD ["postgresql://${DB_USER}:${DB_PASSWORD}@db:5432/${DB_NAME}?sslmode=disable -verbose up"]