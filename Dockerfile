FROM golang:1.22-alpine as dev

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 3000

CMD ["air", "-c", ".air.toml"]