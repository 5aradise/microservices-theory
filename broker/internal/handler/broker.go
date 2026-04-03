package handler

import (
	"micro/broker/internal/model"
	"micro/broker/internal/service"
	contres "micro/common/contracts/http/res"
	util "micro/common/utils"
	"net/http"
)

type Broker struct {
	serv *service.Broker
}

func New(serv *service.Broker) *Broker {
	return new(Broker{
		serv: serv,
	})
}

func (h *Broker) Broker(w http.ResponseWriter, r *http.Request) {
	util.WriteJSON(w, http.StatusAccepted, SubmissionResp{"Hello from the broker", nil})
}

func (h *Broker) Submission(w http.ResponseWriter, r *http.Request) {
	req, err := util.ReadJSON[SubmissionReq](w, r)
	if err != nil {
		contres.WriteError(w, err)
		return
	}

	act, params, err := HttpToActionAndParams(req)
	if err != nil {
		contres.WriteError(w, err)
		return
	}

	data, err := h.serv.Submission(r.Context(), act, params)
	if err != nil {
		contres.WriteError(w, err)
		return
	}

	util.WriteJSON(w, http.StatusAccepted, SubmissionResp{
		Message: "Success!",
		Data:    data,
	})
}

func (h *Broker) GRPCSubmission(w http.ResponseWriter, r *http.Request) {
	req, err := util.ReadJSON[SubmissionReq](w, r)
	if err != nil {
		contres.WriteError(w, err)
		return
	}

	_, params, err := HttpToActionAndParams(req)
	if err != nil {
		contres.WriteError(w, err)
		return
	}

	data, err := h.serv.Submission(r.Context(), model.GRPCLogAction, params)
	if err != nil {
		contres.WriteError(w, err)
		return
	}

	util.WriteJSON(w, http.StatusAccepted, SubmissionResp{
		Message: "Success!",
		Data:    data,
	})
}
