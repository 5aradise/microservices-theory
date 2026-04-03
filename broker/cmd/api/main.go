package main

import (
	"log"
	"micro/broker/internal/external"
	"micro/broker/internal/handler"
	"micro/broker/internal/service"
	amqputil "micro/broker/pkg/amqp"
	"net"
	"net/http"
	"net/rpc"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port        = "80"
	authURL     = "http://authentication-service/authenticate"
	logURL      = "http://logger-service/log"
	mailURL     = "http://mail-service/send"
	rabbitmqURL = "amqp://guest:guest@rabbitmq"
	rpcURL      = "logger-service:5001"
	grpcURL     = "logger-service:50001"
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

	rpcClient, err := rpc.Dial("tcp", rpcURL)
	if err != nil {
		log.Panic(err)
	}

	grpcClient, err := grpc.NewClient(grpcURL,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panic(err)
	}

	authServ := external.NewAuthService(authURL)
	logServ := external.NewLogService(logURL)
	mailServ := external.NewMailService(mailURL)
	queueServ := external.NewQueueService(emitter)
	rpcServ := external.NewRPCService(rpcClient)
	grpcServ := external.NewGRPCService(grpcClient)

	serv := service.NewBroker(authServ, logServ, mailServ, queueServ, rpcServ, grpcServ)

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
