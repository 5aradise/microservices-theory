package main

import (
	"log"
	"micro/broker/internal/external"
	"micro/broker/internal/handler"
	"micro/broker/internal/service"
	"net"
	"net/http"
)

const (
	port    = "80"
	authURL = "http://authentication-service/authenticate"
	logURL  = "http://logger-service/log"
	mailURL = "http://mail-service/send"
)

func main() {
	authServ := external.NewAuthService(authURL)
	logServ := external.NewLogService(logURL)
	mailServ := external.NewMailService(mailURL)

	serv := service.NewBroker(authServ, logServ, mailServ)

	h := handler.New(serv)

	srv := http.Server{
		Addr:    net.JoinHostPort("", port),
		Handler: routes(h),
	}

	log.Println("Starting broker service on port: ", port)

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
