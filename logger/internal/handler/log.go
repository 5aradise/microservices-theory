package handler

import (
	contreq "micro/common/contracts/http/req"
	contres "micro/common/contracts/http/res"
	util "micro/common/utils"
	"micro/logger/internal/model"
	"micro/logger/internal/service"
	"net/http"
)

type Log struct {
	serv *service.Log
}

func NewLog(serv *service.Log) *Log {
	return &Log{
		serv: serv,
	}
}

func (h *Log) WriteLog(w http.ResponseWriter, r *http.Request) {
	req, err := util.ReadJSON[contreq.WriteLog](w, r)
	if err != nil {
		contres.WriteError(w, err)
		return
	}

	err = h.serv.WriteLog(r.Context(), model.Entry{
		Name: req.Name,
		Data: req.Data,
	})
	if err != nil {
		contres.WriteError(w, err)
		return
	}

	util.WriteJSON(w, http.StatusAccepted, contres.WriteLog{
		Message: "logged",
	})
}
