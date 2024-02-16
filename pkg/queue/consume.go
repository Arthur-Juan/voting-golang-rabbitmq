package queue

import (
	"github.com/gofiber/fiber/v2/log"
	amqp "github.com/rabbitmq/amqp091-go"
)

func (r *RabbitMq) Consume(ch *amqp.Channel, out chan amqp.Delivery) error {
	msgs, err := ch.Consume(
		"vote",
		"",
		false, // auto ack
		false, // exclusive
		false, // no local
		false, // no wait
		nil,   //args
	)

	if err != nil {
		log.Errorf("Error => %s", err.Error())
		return err
	}

	for msg := range msgs {
		out <- msg
	}

	return nil
}
