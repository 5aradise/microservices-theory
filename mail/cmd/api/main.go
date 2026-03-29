package main

import (
	"log"
	"micro/mail/internal/handler"
	"micro/mail/internal/service"
	"net"
	"net/http"
)

const port = "80"

func main() {
	s, err := service.NewMail(service.MailParamsFromEnv())
	if err != nil {
		log.Panic(err)
	}

	h := handler.NewMail(s)

	srv := http.Server{
		Addr:    net.JoinHostPort("", port),
		Handler: routes(h),
	}

	log.Println("Starting mail service on port: ", port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
