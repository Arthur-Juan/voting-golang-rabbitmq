package queue

import (
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMq struct {
	ConnStr    string
	Connection *amqp.Channel
}

func NewQueue() *RabbitMq {
	return &RabbitMq{
		ConnStr: os.Getenv("QUEUE_URL"),
	}
}

func (r *RabbitMq) Connect(connstring string) (*amqp.Channel, error) {
	conn, err := amqp.Dial(r.ConnStr)

	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return ch, nil
}

func (r *RabbitMq) Disconnect(ch *amqp.Channel) {
	ch.Close()
}
