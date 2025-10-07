// internal/config/rabbitmq_config.go
package config

import (
	"log"
	"sync"
	"time"
	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	RabbitMQConn    *amqp.Connection
	RabbitMQChannel *amqp.Channel
	once            sync.Once
	mu              sync.Mutex
)

// connect dials and opens a channel. The caller is responsible for closing
// the connection on failure.
func connect(rabbitmqURL string) (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial(rabbitmqURL)
	if err != nil {
		return nil, nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		_ = conn.Close()
		return nil, nil, err
	}

	return conn, ch, nil
}

// InitializeRabbitMQ establishes a connection/channel and starts a background
// goroutine to monitor connection closures and reconnect automatically.
func InitializeRabbitMQ(rabbitmqURL string) {
	once.Do(func() {
		// initial connect (block until success)
		for {
			conn, ch, err := connect(rabbitmqURL)
			if err != nil {
				log.Printf("Failed to connect to RabbitMQ: %v; retrying in 5s", err)
				time.Sleep(5 * time.Second)
				continue
			}

			mu.Lock()
			RabbitMQConn = conn
			RabbitMQChannel = ch
			mu.Unlock()

			log.Println("Connected to RabbitMQ")
			break
		}

		// Watch the connection for closure and reconnect in a loop
		go func(url string) {
			for {
				// capture the current connection to watch its close channel
				mu.Lock()
				connToWatch := RabbitMQConn
				mu.Unlock()

				if connToWatch == nil {
					// if somehow nil, try to reconnect immediately
					time.Sleep(2 * time.Second)
				} else {
					// Block until the connection is closed
					err := <-connToWatch.NotifyClose(make(chan *amqp.Error))
					if err != nil {
						log.Printf("RabbitMQ connection closed: %v", err)
					} else {
						log.Println("RabbitMQ connection closed")
					}
				}

				// Attempt reconnect with backoff
				for {
					conn, ch, err := connect(url)
					if err != nil {
						log.Printf("Failed to reconnect to RabbitMQ: %v; retrying in 5s", err)
						time.Sleep(5 * time.Second)
						continue
					}

					mu.Lock()
					RabbitMQConn = conn
					RabbitMQChannel = ch
					mu.Unlock()

					log.Println("Reconnected to RabbitMQ")
					break
				}
			}
		}(rabbitmqURL)
	})
}

func CloseRabbitMQ() {
	mu.Lock()
	defer mu.Unlock()
	if RabbitMQChannel != nil {
		_ = RabbitMQChannel.Close()
		RabbitMQChannel = nil
	}
	if RabbitMQConn != nil {
		_ = RabbitMQConn.Close()
		RabbitMQConn = nil
	}
	log.Println("RabbitMQ connection closed")
}
