package queue

import (
	"log"
	"strconv"
	"time"

	"github.com/dipo0x/golang-url-shortener/internal/config"
	"github.com/dipo0x/golang-url-shortener/internal/infra"
	amqp "github.com/rabbitmq/amqp091-go"
)

func PublishJob(urlID string, hours float64) {
	if config.RabbitMQChannel == nil {
		log.Fatal("RabbitMQ channel is not initialized")
	}
	delayQueueArgs := amqp.Table{
		"x-dead-letter-exchange": "delayed-exchange",
		"x-dead-letter-routing-key": "job_key",
	}

	_, err := config.RabbitMQChannel.QueueDeclare(
		"delay_queue",
		true,
		false,
		false,
		false,
		delayQueueArgs,
	)
	infra.FailOnError(err, "Failed to declare delay queue")

	duration := time.Duration(hours * float64(time.Hour))
	expiration := strconv.FormatInt(duration.Milliseconds(), 10)

	body := "deleteOldUrl:" + urlID

	err = config.RabbitMQChannel.Publish(
		"", // default exchange routes to queue by name
		"delay_queue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
			Expiration:  expiration,
		},
	)
	infra.FailOnError(err, "Failed to publish job")

	log.Printf(" [x] Scheduled job with URL '%s' to run in %v hours", urlID, hours)
}
