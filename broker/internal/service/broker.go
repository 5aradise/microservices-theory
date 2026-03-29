package service

import (
	"context"
	"errors"
	"micro/broker/internal/external"
	"micro/broker/internal/model"
)

type Broker struct {
	authServ *external.AuthService
	logServ  *external.LogService
	mailServ *external.MailService
}

func NewBroker(authServ *external.AuthService, logServ *external.LogService, mailServ *external.MailService) *Broker {
	return new(Broker{
		authServ: authServ,
		logServ:  logServ,
		mailServ: mailServ,
	})
}

func (s *Broker) Submission(ctx context.Context, action model.Action, params model.SubmissionParams) (data any, err error) {
	switch action {
	case model.AuthAction:
		data, err = s.authServ.Authenticate(ctx, *params.Auth)
	case model.LogAction:
		data, err = s.logServ.Log(ctx, *params.Log)
	case model.MailAction:
		data, err = s.mailServ.SendMail(ctx, *params.Mail)
	default:
		panic(errors.New("invalid action"))
	}
	return
}
