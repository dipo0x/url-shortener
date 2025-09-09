package queue

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/dipo0x/golang-url-shortener/internal/config"
	"github.com/dipo0x/golang-url-shortener/internal/infra"
)

var ctx = context.Background()

func deleteURL(urlID string) {
	_, err := config.Pool.Exec(ctx, `DELETE FROM urls WHERE id = $1`, urlID)
	if err != nil {
		log.Printf("Failed to delete URL %s: %v", urlID, err)
		return
	}
	log.Printf("URL deleted: %s", urlID)
}

func StartConsumer() {
	ch := config.RabbitMQChannel

	args := amqp.Table{"x-delayed-type": "direct"}
	err := ch.ExchangeDeclare(
		"delayed-exchange", 
		"x-delayed-message", 
		true,  
		false,  
		false, 
		false,  
		args,   
	)
	infra.FailOnError(err, "Failed to declare exchange")

	q, err := ch.QueueDeclare("jobs_queue", true, false, false, false, nil)
	infra.FailOnError(err, "Failed to declare queue")

	err = ch.QueueBind(q.Name, "job_key", "delayed-exchange", false, nil)
	infra.FailOnError(err, "Failed to bind queue")

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	infra.FailOnError(err, "Failed to register consumer")

	go func() {
		for d := range msgs {
			parts := strings.Split(string(d.Body), ":")
			if len(parts) < 2 {
				log.Printf("Invalid message format: %s", d.Body)
				continue
			}

			fn, arg := parts[0], parts[1]
			switch fn {
			case "deleteOldUrl":
				deleteURL(arg)
			default:
				log.Printf("Unknown function: %s", fn)
			}
		}
	}()

	log.Printf("Waiting for jobs...")

	// Graceful shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	log.Println("Shutting down consumer...")
}
