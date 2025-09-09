// internal/config/rabbitmq_config.go
package config

import (
	"log"
	"sync"
	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	RabbitMQConn    *amqp.Connection
	RabbitMQChannel *amqp.Channel
	once            sync.Once
)

func InitializeRabbitMQ(rabbitmqURL string) {
	once.Do(func() {
		conn, err := amqp.Dial(rabbitmqURL)
		if err != nil {
			log.Fatalf("Failed to connect to RabbitMQ: %v", err)
		}

		ch, err := conn.Channel()
		if err != nil {
			log.Fatalf("Failed to open a channel: %v", err)
		}

		log.Println("Connected to RabbitMQ")

		RabbitMQConn = conn
		RabbitMQChannel = ch
	})
}

func CloseRabbitMQ() {
	if RabbitMQChannel != nil {
		_ = RabbitMQChannel.Close()
	}
	if RabbitMQConn != nil {
		_ = RabbitMQConn.Close()
	}
	log.Println("RabbitMQ connection closed")
}
