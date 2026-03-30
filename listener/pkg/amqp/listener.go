package amqputil

import amqp "github.com/rabbitmq/amqp091-go"

type Listener struct {
	ch           *amqp.Channel
	exchangeName string
}

func NewListener(ch *amqp.Channel, exchangeName string) (*Listener, error) {
	err := declareExchange(ch,
		exchangeName,
		"topic",
	)
	if err != nil {
		return nil, err
	}

	return new(Listener{
		ch:           ch,
		exchangeName: exchangeName,
	}), nil
}

func (l *Listener) Consume(topics []string, handleMessage func(msg amqp.Delivery)) error {
	q, err := declareRandomQueue(l.ch)
	if err != nil {
		return err
	}

	for _, topic := range topics {
		err := l.ch.QueueBind(
			q.Name,
			topic,
			l.exchangeName,
			false, nil)

		if err != nil {
			return err
		}
	}

	msgs, err := l.ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	for msg := range msgs {
		go handleMessage(msg)
	}

	return nil
}
