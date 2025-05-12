package helpers

import (
	"go.mongodb.org/mongo-driver/bson"
	"crypto/rand"
	"log"
	"encoding/hex"
	"context"
	"github.com/dipo0x/golang-url-shortener/internal/config"
)

func generateRandomString() string {
	bytes := make([]byte, 5) // 5 bytes = 10 hex characters
	_, err := rand.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(bytes)
}

func GetUniqueRandomString() (string, error) {
	for {
		randomString := generateRandomString()

		count, err := config.MongoDatabase.Collection("urls").CountDocuments(context.TODO(), bson.M{"randomField": randomString})
		if err != nil {
			return "", err
		}

		if count == 0 {
			return randomString, nil
		}
	}
}
