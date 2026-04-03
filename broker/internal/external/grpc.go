package external

import (
	"context"
	"micro/broker/internal/logs"
	"micro/broker/internal/model"

	"google.golang.org/grpc"
)

type GRPCService struct {
	client logs.LogServiceClient
}

func NewGRPCService(client grpc.ClientConnInterface) *GRPCService {
	return new(GRPCService{
		client: logs.NewLogServiceClient(client),
	})
}

func (s *GRPCService) Log(ctx context.Context, params model.LogParams) (data any, err error) {
	res, err := s.client.WriteLog(ctx, &logs.LogRequest{
		LogEntry: &logs.Log{
			Name: params.Name,
			Data: params.Data,
		},
	})
	if err != nil {
		return nil, err
	}

	return res.Result, nil
}
