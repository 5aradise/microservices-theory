package main

import (
	"log"
	"micro/listener/internal/external"
	"micro/listener/internal/handler"
	amqputil "micro/listener/pkg/amqp"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	rabbitmqURL = "amqp://guest:guest@rabbitmq"
	logURL      = "http://logger-service/log"
)

func main() {
	conn, err := amqp.Dial(rabbitmqURL)
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Panic(err)
	}
	defer ch.Close()

	l, err := amqputil.NewListener(ch, "logs")
	if err != nil {
		log.Panic(err)
	}

	logServ := external.NewLogService(logURL)

	h := handler.New(logServ)

	err = l.Consume([]string{"log.INFO", "log.WARNING", "log.ERROR"}, h.Handle)
	if err != nil {
		log.Panic(err)
	}
}
