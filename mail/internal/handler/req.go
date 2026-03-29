package handler

import (
	contreq "micro/common/contracts/http/req"
	"micro/mail/internal/model"
)

func HttpToMessage(req contreq.SendMail) model.Message {
	return model.Message{
		From:    req.From,
		To:      req.To,
		Subject: req.Subject,
		Body:    req.Body,
	}
}
