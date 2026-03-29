package main

import (
	"context"
	"log"
	"micro/authentication/internal/external"
	"micro/authentication/internal/handler"
	"micro/authentication/internal/service"
	storage "micro/authentication/internal/storage/postgresql"
	"micro/authentication/pkg/postgresql"
	"net"
	"net/http"
)

const (
	port   = "80"
	logURL = "http://logger-service/log"
)

func main() {
	db, err := postgresql.New(context.TODO(), postgresql.DSNFromEnv())
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	stor := storage.NewUsers(db)

	logServ := external.NewLogService(logURL)

	serv := service.NewAuth(stor, logServ)

	h := handler.NewAuth(serv)

	srv := http.Server{
		Addr:    net.JoinHostPort("", port),
		Handler: routes(h),
	}

	log.Println("Starting authentication service on port: ", port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
