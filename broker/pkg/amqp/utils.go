package amqputil

import amqp "github.com/rabbitmq/amqp091-go"

func declareExchange(ch *amqp.Channel, name, kind string) error {
	return ch.ExchangeDeclare(
		name, kind,
		true,
		false,
		false,
		false,
		nil,
	)
}

func declareRandomQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
}
