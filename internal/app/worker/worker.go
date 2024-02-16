package worker

import (
	"fmt"
	"os"
	"time"

	"github.com/arthur-juan/voting-golang-rabbitmq/pkg/queue"
	"github.com/gofiber/fiber/v2/log"
	amqp "github.com/rabbitmq/amqp091-go"
)

func Run(start time.Time) {
	log.Debug("Running")

	duration := time.Until(start.Add(3 * time.Hour))

	log.Debug(duration)
	if duration > 0 {
		time.Sleep(duration)
	} else {
		log.Debug("Start time is in the past. Not sleeping.")
	}

	queue := queue.NewQueue()
	ch, err := queue.Connect(queue.ConnStr)

	if err != nil {
		panic(err)
	}

	defer queue.Disconnect(ch)
	msg := make(chan amqp.Delivery)
	go queue.Consume(ch, msg)

	counter := os.Getenv("WORKER_COUNTER")

	for i := 0; i < len(counter); i++ {
		go CountVotes(msg)
	}

}

func CountVotes(data <-chan amqp.Delivery) {
	for msg := range data {
		fmt.Println(msg)
	}
}
