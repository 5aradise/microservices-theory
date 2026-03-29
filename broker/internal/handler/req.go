package handler

import (
	"errors"
	"micro/broker/internal/model"
)

type SubmissionReq struct {
	Action string   `json:"action"`
	Auth   *AuthReq `json:"auth"`
	Log    *LogReq  `json:"log"`
	Mail   *MailReq `json:"mail"`
}

type AuthReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LogReq struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type MailReq struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

func HttpToActionAndParams(req SubmissionReq) (model.Action, model.SubmissionParams, error) {
	switch req.Action {
	case "auth":
		if req.Auth == nil {
			return 0, model.SubmissionParams{}, errors.New("invalid params")
		}
		return model.AuthAction, model.SubmissionParams{
			Auth: &model.AuthParams{
				Email:    req.Auth.Email,
				Password: req.Auth.Password,
			},
		}, nil
	case "log":
		if req.Log == nil {
			return 0, model.SubmissionParams{}, errors.New("invalid params")
		}
		return model.LogAction, model.SubmissionParams{
			Log: &model.LogParams{
				Name: req.Log.Name,
				Data: req.Log.Data,
			},
		}, nil
	case "mail":
		if req.Mail == nil {
			return 0, model.SubmissionParams{}, errors.New("invalid params")
		}
		return model.MailAction, model.SubmissionParams{
			Mail: &model.MailParams{
				From:    req.Mail.From,
				To:      req.Mail.To,
				Subject: req.Mail.Subject,
				Message: req.Mail.Message,
			},
		}, nil
	default:
		return 0, model.SubmissionParams{}, errors.New("invalid action")
	}
}
