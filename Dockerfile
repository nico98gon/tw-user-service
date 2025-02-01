FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/air-verse/air@latest

COPY . .

RUN go env
RUN go build -v -o user-service ./cmd/main.go

# FROM debian:buster-slim
FROM golang:1.23-alpine

WORKDIR /app

COPY --from=builder /app/user-service .
COPY --from=builder /go/bin/air /usr/local/bin/air

# COPY wait-for-it.sh /wait-for-it.sh
# RUN chmod +x /wait-for-it.sh
RUN chmod +x ./user-service

EXPOSE 8081

CMD ["/wait-for-it.sh", "postgres:5432", "--", "air", "-c", ".air.toml"]
