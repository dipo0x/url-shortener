package main

import (
	"context"
	"encoding/json"
	"log"
	"github.com/dipo0x/golang-url-shortener/config"
	"github.com/dipo0x/golang-url-shortener/types"
	"github.com/hibiken/asynq"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/google/uuid"
)

const (
	TypeDeleteDocument = "mongo:delete_document"
)

func main() {
	if err := config.InitializeMongoDB(config.Config("MONGO_URI"), config.Config("MONGO_DATABASE")); err != nil {
		log.Fatalf("MongoDB connection failed: %v", err)
	}
	
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: config.Config("REDIS_URL")},
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				"default": 6,
				"critical": 4,
			},
		},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(TypeDeleteDocument, func(ctx context.Context, task *asynq.Task) error {
		var p types.IDeletePayload
		if err := json.Unmarshal(task.Payload(), &p); err != nil {
			return err
		}
		objID, _ := uuid.Parse(p.DocumentID)
		filter := bson.M{"_id": objID}
		log.Println(objID, filter)
		_, err := config.MongoDatabase.Collection("urls").DeleteOne(context.TODO(), filter)
		if err != nil {
			log.Println("Delete error:", err)
			return err
		}
		return nil
	})
	if err := srv.Run(mux); err != nil {
		log.Fatal(err)
	}
}
