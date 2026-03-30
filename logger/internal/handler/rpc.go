package handler

import (
	"context"
	contreq "micro/common/contracts/http/req"
	contres "micro/common/contracts/http/res"
	"micro/logger/internal/model"
	"micro/logger/internal/service"
)

type RPCLog struct {
	serv *service.Log
}

func NewRPCLog(serv *service.Log) *RPCLog {
	return &RPCLog{
		serv: serv,
	}
}

func (h *RPCLog) WriteLog(req contreq.WriteLog, res *contres.WriteLog) error {
	err := h.serv.WriteLog(context.TODO(), model.Entry{
		Name: req.Name,
		Data: req.Data,
	})
	if err != nil {
		return err
	}

	res.Message = "logged via rpc"
	return nil
}
