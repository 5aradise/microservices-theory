# Build the binary
FROM golang:1.26-alpine AS builder

RUN mkdir /app

COPY .. /app

WORKDIR /app

RUN go build -o broker-app ./broker/cmd/api

RUN chmod +x /app/broker-app

# Build a small image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/broker-app /app

CMD ["/app/broker-app"]
