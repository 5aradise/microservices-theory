package main

import (
	"context"
	"log"
	"micro/logger/internal/handler"
	"micro/logger/internal/service"
	storage "micro/logger/internal/storage/mongodb"
	"micro/logger/pkg/mongodb"
	"net"
	"net/http"
	"net/rpc"
)

const (
	port     = "80"
	rpcPort  = "5001"
	grpcPort = "50001"
)

func main() {
	mongoClient, err := mongodb.New(context.TODO(), mongodb.DSNFromEnv())
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		if err := mongoClient.Disconnect(context.TODO()); err != nil {
			log.Panic(err)
		}
	}()

	stor := storage.NewEntries(mongoClient)

	serv := service.NewLog(stor)

	h := handler.NewLog(serv)

	rpc.Register(handler.NewRPCLog(serv))

	l, err := net.Listen("tcp", net.JoinHostPort("", rpcPort))
	if err != nil {
		log.Panic(err)
	}
	defer l.Close()

	log.Println("Starting rpc logger service on port: ", rpcPort)

	srv := http.Server{
		Addr:    net.JoinHostPort("", port),
		Handler: routes(h),
	}

	log.Println("Starting logger service on port: ", port)

	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				break
			}

			go rpc.ServeConn(conn)
		}
	}()

	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
