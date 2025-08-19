FROM golang:1.25.0-bookworm AS deps

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

FROM golang:1.25.0-bookworm AS builder

WORKDIR /app

COPY --from=deps /go/pkg /go/pkg
COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux

RUN go build -ldflags="-w -s" -o mep-api-core .

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/mep-api-core .
COPY bin/start_server.sh bin/start_server.sh

RUN chmod +x /app/bin/start_server.sh

EXPOSE 8080

CMD ["bin/start_server.sh"]
