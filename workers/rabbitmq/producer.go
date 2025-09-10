package queue

import (
	"log"
	"time"

	"github.com/dipo0x/golang-url-shortener/internal/config"
	"github.com/dipo0x/golang-url-shortener/internal/infra"
	amqp "github.com/rabbitmq/amqp091-go"
)

func PublishJob(urlID string, hours float64) {
	if config.RabbitMQChannel == nil {
		log.Fatal("RabbitMQ channel is not initialized")
	}

	duration := time.Duration(hours * float64(time.Hour))
	headers := amqp.Table{"x-delay": duration.Milliseconds()}

	body := "deleteOldUrl:" + urlID

	err := config.RabbitMQChannel.Publish(
		"delayed-exchange",
		"job_key", // in your project, you can rename this. just ensure it matches what y've in your consumer.go.
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
			Headers:     headers,
		},
	)
	infra.FailOnError(err, "Failed to publish job")

	log.Printf(" [x] Scheduled job with URL '%s' to run in %v hours", urlID, hours)
}
