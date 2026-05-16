FROM golang:latest AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/ethchess-bot-tewdros .


FROM debian:stable-slim

RUN apt-get update && apt-get install -y \
    stockfish \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY --from=builder /app/bin/ethchess-bot-tewdros ./bin/ethchess-bot-tewdros

CMD ["./bin/ethchess-bot-tewdros"]
