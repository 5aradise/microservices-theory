package main

import (
	"log"
	"micro/broker/internal/external"
	"micro/broker/internal/handler"
	"micro/broker/internal/service"
	amqputil "micro/broker/pkg/amqp"
	"net"
	"net/http"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	port        = "80"
	authURL     = "http://authentication-service/authenticate"
	logURL      = "http://logger-service/log"
	mailURL     = "http://mail-service/send"
	rabbitmqURL = "amqp://guest:guest@rabbitmq"
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

	emitter, err := amqputil.NewEmitter(ch, "logs")
	if err != nil {
		log.Panic(err)
	}

	authServ := external.NewAuthService(authURL)
	logServ := external.NewLogService(logURL)
	mailServ := external.NewMailService(mailURL)
	queueServ := external.NewQueueService(emitter)

	serv := service.NewBroker(authServ, logServ, mailServ, queueServ)

	h := handler.New(serv)

	srv := http.Server{
		Addr:    net.JoinHostPort("", port),
		Handler: routes(h),
	}

	log.Println("Starting broker service on port: ", port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
