package helpers

import (
	"crypto/rand"
	"log"
	"encoding/hex"
	"context"
	"github.com/dipo0x/golang-url-shortener/internal/config"
)

func generateRandomString() string {
	bytes := make([]byte, 5)
	_, err := rand.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(bytes)
}

func GetUniqueRandomString() (string, error) {
	for {
		var count int
		randomString := generateRandomString()

		err := config.Pool.QueryRow(
		context.Background(),
		`SELECT COUNT(*) FROM urls WHERE short_url = $1`,
		randomString,
	).Scan(&count)
		if err != nil {
			return "", err
		}

		if count == 0 {
			return randomString, nil
		}
	}
}
