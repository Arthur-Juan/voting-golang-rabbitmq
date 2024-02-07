package queue

import (
	"context"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (r *RabbitMq) Publish(ch *amqp.Channel, message string) error {

	queue, err := ch.QueueDeclare(
		"vote",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = ch.PublishWithContext(ctx,
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})

	if err != nil {
		return err
	}

	return nil
}
