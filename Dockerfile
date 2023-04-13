## --- Builder image --- ##
FROM golang:1.20.3-bullseye AS builder

WORKDIR /cefetdb2

COPY ./backend/go.* ./

COPY ./backend ./

RUN go build -o api ./cmd/main

RUN go build -o get_oauth_token ./cmd/auth

## --- Runner image --- ##
FROM debian:bullseye-slim

WORKDIR /cefetdb2

RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /cefetdb2/api /cefetdb2/get_oauth_token ./

COPY credentials.json config.yaml token.json ./

EXPOSE 3500

CMD ["./api"]