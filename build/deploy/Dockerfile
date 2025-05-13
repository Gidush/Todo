FROM golang:1.23-alpine AS builder
WORKDIR /app
EXPOSE 8080

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o=/app/bin ./cmd/todo/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/bin /app/bin
COPY --from=builder /app/build/env /app/build/env
ENTRYPOINT ["/app/bin/main"]