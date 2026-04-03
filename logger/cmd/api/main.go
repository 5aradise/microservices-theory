package main

import (
	"context"
	"log"
	"micro/logger/internal/handler"
	"micro/logger/internal/logs"
	"micro/logger/internal/service"
	storage "micro/logger/internal/storage/mongodb"
	"micro/logger/pkg/mongodb"
	"net"
	"net/http"
	"net/rpc"

	"google.golang.org/grpc"
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

	grpcs := grpc.NewServer()
	grpch := handler.NewGRPC(serv)
	logs.RegisterLogServiceServer(grpcs, grpch)

	rpcl, err := net.Listen("tcp", net.JoinHostPort("", rpcPort))
	if err != nil {
		log.Panic(err)
	}
	defer rpcl.Close()

	grpcl, err := net.Listen("tcp", net.JoinHostPort("", grpcPort))
	if err != nil {
		log.Panic(err)
	}
	defer grpcl.Close()

	log.Println("Starting rpc logger service on port: ", rpcPort)

	srv := http.Server{
		Addr:    net.JoinHostPort("", port),
		Handler: routes(h),
	}

	log.Println("Starting logger service on port: ", port)

	go func() {
		for {
			conn, err := rpcl.Accept()
			if err != nil {
				break
			}

			go rpc.ServeConn(conn)
		}
	}()

	go func() {
		err = grpcs.Serve(grpcl)
		if err != nil {
			log.Panic(err)
		}
	}()

	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
