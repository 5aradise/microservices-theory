package external

import (
	"micro/broker/internal/model"
	contreq "micro/common/contracts/http/req"
)

func AuthToReq(auth model.AuthParams) contreq.Authenticate {
	return contreq.Authenticate{
		Email:    auth.Email,
		Password: auth.Password,
	}
}

func LogToReq(log model.LogParams) contreq.WriteLog {
	return contreq.WriteLog{
		Name: log.Name,
		Data: log.Data,
	}
}

func MailToReq(mail model.MailParams) contreq.SendMail {
	return contreq.SendMail{
		From:    mail.From,
		To:      []string{mail.To},
		Subject: mail.Subject,
		Body:    mail.Message,
	}
}
