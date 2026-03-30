# Build the binary
FROM golang:1.26-alpine AS builder

RUN mkdir /app

COPY .. /app

WORKDIR /app

RUN go build -o listener-app ./listener/cmd

RUN chmod +x /app/listener-app

# Build a small image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/listener-app /app

CMD ["/app/listener-app"]
