# Build the binary
FROM golang:1.26-alpine AS builder

RUN mkdir /app

COPY .. /app

WORKDIR /app

RUN go build -o logger-app ./logger/cmd/api

RUN chmod +x /app/logger-app

# Build a small image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/logger-app /app

CMD ["/app/logger-app"]
