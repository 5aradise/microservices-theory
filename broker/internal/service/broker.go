package service

import (
	"context"
	"micro/broker/internal/external"
	"micro/broker/internal/model"
)

type Broker struct {
	authServ  *external.AuthService
	logServ   *external.LogService
	mailServ  *external.MailService
	queueServ *external.QueueService
	rpcServ   *external.RPCService
	grpcServ  *external.GRPCService
}

func NewBroker(authServ *external.AuthService, logServ *external.LogService,
	mailServ *external.MailService, queueServ *external.QueueService,
	rpcServ *external.RPCService, grpcServ *external.GRPCService) *Broker {
	return new(Broker{
		authServ:  authServ,
		logServ:   logServ,
		mailServ:  mailServ,
		queueServ: queueServ,
		rpcServ:   rpcServ,
		grpcServ:  grpcServ,
	})
}

func (s *Broker) Submission(ctx context.Context, action model.Action, params model.SubmissionParams) (data any, err error) {
	switch action {
	case model.AuthAction:
		data, err = s.authServ.Authenticate(ctx, *params.Auth)
	case model.GRPCLogAction:
		data, err = s.grpcServ.Log(ctx, *params.Log)
	case model.LogAction:
		data, err = s.rpcServ.Log(ctx, *params.Log)
	case model.MailAction:
		data, err = s.mailServ.SendMail(ctx, *params.Mail)
	default:
		panic("invalid action")
	}
	return
}
