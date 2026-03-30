package external

import (
	"context"
	"encoding/json"
	"micro/broker/internal/model"
	amqputil "micro/broker/pkg/amqp"
)

type QueueService struct {
	emitter *amqputil.Emitter
}

func NewQueueService(emitter *amqputil.Emitter) *QueueService {
	return new(QueueService{
		emitter: emitter,
	})
}

func (s *QueueService) Log(ctx context.Context, params model.QueueParams) error {
	logBody, err := json.Marshal(LogToReq(params.Data))
	if err != nil {
		return err
	}

	return s.emitter.Push(params.Key, logBody)
}
