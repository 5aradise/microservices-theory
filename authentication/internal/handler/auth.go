package handler

import (
	"micro/authentication/internal/service"
	contreq "micro/common/contracts/http/req"
	contres "micro/common/contracts/http/res"
	util "micro/common/utils"
	"net/http"
)

type Auth struct {
	serv *service.Auth
}

func NewAuth(serv *service.Auth) *Auth {
	return &Auth{
		serv: serv,
	}
}

func (h *Auth) Authenticate(w http.ResponseWriter, r *http.Request) {
	body, err := util.ReadJSON[contreq.Authenticate](w, r)
	if err != nil {
		contres.WriteError(w, err)
		return
	}

	user, err := h.serv.Authenticate(r.Context(), body.Email, body.Password)
	if err != nil {
		contres.WriteError(w, err, http.StatusUnauthorized)
		return
	}

	util.WriteJSON(w, http.StatusAccepted, UserToResp(user))
}
