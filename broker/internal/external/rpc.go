package external

import (
	"context"
	"micro/broker/internal/model"
	contres "micro/common/contracts/http/res"
	"net/rpc"
)

type RPCService struct {
	client *rpc.Client
}

func NewRPCService(client *rpc.Client) *RPCService {
	return new(RPCService{
		client: client,
	})
}

func (s *RPCService) Log(ctx context.Context, params model.LogParams) (data any, err error) {
	var res contres.WriteLog
	err = s.client.Call("RPCLog.WriteLog", LogToReq(params), &res)
	if err != nil {
		return nil, err
	}

	return res.Message, nil
}
