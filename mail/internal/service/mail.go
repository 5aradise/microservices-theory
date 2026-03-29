package service

import (
	"micro/mail/internal/model"
	"net"
	"net/smtp"
	"os"
)

type Mail struct {
	addr string
	auth smtp.Auth
}

type MailParams struct {
	Host     string
	Port     string
	Username string
	Password string
}

func MailParamsFromEnv() MailParams {
	return MailParams{
		Host:     os.Getenv("MAIL_HOST"),
		Port:     os.Getenv("MAIL_PORT"),
		Username: os.Getenv("MAIL_USERNAME"),
		Password: os.Getenv("MAIL_PASSWORD"),
	}
}

func NewMail(params MailParams) (*Mail, error) {
	auth := smtp.PlainAuth("", params.Username, params.Password, params.Host)
	mail := new(Mail{
		addr: net.JoinHostPort(params.Host, params.Port),
		auth: auth,
	})
	if err := mail.ping(); err != nil {
		return nil, err
	}
	return mail, nil
}

func (s *Mail) ping() error {
	c, err := smtp.Dial(s.addr)
	if err != nil {
		return err
	}
	defer c.Close()

	return c.Noop()
}

func (s *Mail) SendMail(m model.Message) error {
	return smtp.SendMail(s.addr, s.auth, m.From, m.To, m.Raw())
}
