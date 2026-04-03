package handler

import (
	"context"
	"micro/logger/internal/logs"
	"micro/logger/internal/model"
	"micro/logger/internal/service"
)

type grpcHandler struct {
	logs.UnimplementedLogServiceServer
	serv *service.Log
}

func NewGRPC(serv *service.Log) logs.LogServiceServer {
	return &grpcHandler{
		serv: serv,
	}
}

func (h *grpcHandler) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error) {
	input := req.GetLogEntry()

	err := h.serv.WriteLog(ctx, model.Entry{
		Name: input.Name,
		Data: input.Data,
	})
	if err != nil {
		return nil, err
	}

	return &logs.LogResponse{Result: "logged via grpc"}, nil
}
