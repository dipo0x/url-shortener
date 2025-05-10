package config

import (
		"github.com/hibiken/asynq"
		"github.com/redis/go-redis/v9"
		"log"
		"context"
)

var AsynqClient *asynq.Client

func InitializeRedis(redisURL string) *asynq.Client {
	opt := asynq.RedisClientOpt{Addr: redisURL}
	
	rdb := redis.NewClient(&redis.Options{
		Addr: redisURL,
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis at %s: %v", redisURL, err)
	}
	log.Printf("Connected to Redis Successfully")
	AsynqClient = asynq.NewClient(opt)
	return AsynqClient
}