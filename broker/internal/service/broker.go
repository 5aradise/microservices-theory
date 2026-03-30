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
}

func NewBroker(authServ *external.AuthService, logServ *external.LogService,
	mailServ *external.MailService, queueServ *external.QueueService) *Broker {
	return new(Broker{
		authServ:  authServ,
		logServ:   logServ,
		mailServ:  mailServ,
		queueServ: queueServ,
	})
}

func (s *Broker) Submission(ctx context.Context, action model.Action, params model.SubmissionParams) (data any, err error) {
	switch action {
	case model.AuthAction:
		data, err = s.authServ.Authenticate(ctx, *params.Auth)
	case model.LogAction:
		err = s.queueServ.Emit(ctx, model.QueueParams{
			Key:  "log.INFO",
			Data: *params.Log,
		})
	case model.MailAction:
		data, err = s.mailServ.SendMail(ctx, *params.Mail)
	default:
		panic("invalid action")
	}
	return
}
