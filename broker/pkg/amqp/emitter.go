package amqputil

import amqp "github.com/rabbitmq/amqp091-go"

type Emitter struct {
	ch           *amqp.Channel
	exchangeName string
}

func NewEmitter(ch *amqp.Channel, exchangeName string) (*Emitter, error) {
	err := declareExchange(ch,
		exchangeName,
		"topic",
	)
	if err != nil {
		return nil, err
	}

	return new(Emitter{
		ch:           ch,
		exchangeName: exchangeName,
	}), nil
}

func (e *Emitter) Push(key string, data []byte) error {
	return e.ch.Publish(
		e.exchangeName,
		key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		},
	)
}
