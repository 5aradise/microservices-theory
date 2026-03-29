package handler

import (
	"fmt"
	contreq "micro/common/contracts/http/req"
	contres "micro/common/contracts/http/res"
	util "micro/common/utils"
	"micro/mail/internal/service"
	"net/http"
)

type Mail struct {
	mail *service.Mail
}

func NewMail(mail *service.Mail) *Mail {
	return &Mail{
		mail: mail,
	}
}

func (h *Mail) SendMail(w http.ResponseWriter, r *http.Request) {
	req, err := util.ReadJSON[contreq.SendMail](w, r)
	if err != nil {
		contres.WriteError(w, err)
		return
	}

	err = h.mail.SendMail(HttpToMessage(req))
	if err != nil {
		contres.WriteError(w, err)
		return
	}

	util.WriteJSON(w, http.StatusAccepted, contres.SendMail{
		Message: fmt.Sprintf("sent to %s", req.To),
	})
}
