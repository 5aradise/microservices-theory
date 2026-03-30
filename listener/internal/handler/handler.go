package handler

import (
	"context"
	"encoding/json"
	"log"
	contreq "micro/common/contracts/http/req"
	"micro/listener/internal/external"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Handler struct {
	logServ *external.LogService
}

func New(logServ *external.LogService) *Handler {
	return &Handler{
		logServ: logServ,
	}
}

func (h *Handler) Handle(msg amqp.Delivery) {
	var payload contreq.WriteLog
	err := json.Unmarshal(msg.Body, &payload)
	if err != nil {
		log.Println(err)
		return
	}

	switch payload.Name {
	case "log", "event":
		_, err := h.logServ.Log(context.TODO(), payload)
		if err != nil {
			log.Println(err)
		}

	case "auth":
		//TODO

	default:
		_, err := h.logServ.Log(context.TODO(), payload)
		if err != nil {
			log.Println(err)
		}
	}
}
