# Build the binary
FROM golang:1.26-alpine AS builder

RUN mkdir /app

COPY .. /app

WORKDIR /app

RUN go build -o auth-app ./authentication/cmd/api

RUN chmod +x /app/auth-app

# Build a small image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/auth-app /app

CMD ["/app/auth-app"]
