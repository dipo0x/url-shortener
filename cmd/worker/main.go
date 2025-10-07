// cmd/worker/main.go
package main

import (
	"log"

	"github.com/dipo0x/golang-url-shortener/internal/config"
	"github.com/dipo0x/golang-url-shortener/workers/rabbitmq"
)

func main() {
	config.InitializeRabbitMQ(config.Config("RABBIT_MQ_URL"))

	if err := config.InitializeDB(config.Config("DATABASE_URL")); err != nil {
		log.Fatalf("DB init failed: %v", err)
	}
	defer config.DisconnectDB()

	queue.StartConsumer()
}
