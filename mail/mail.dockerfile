# Build the binary
FROM golang:1.26-alpine AS builder

RUN mkdir /app

COPY .. /app

WORKDIR /app

RUN go build -o mail-app ./mail/cmd/api

RUN chmod +x /app/mail-app

# Build a small image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/mail-app /app

CMD ["/app/mail-app"]
