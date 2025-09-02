package config

import (
	"context"
    "fmt"
    "log"
	"github.com/jackc/pgx/v5/pgxpool"

)

var Pool *pgxpool.Pool
var ctx = context.Background()

func InitializeDB(uri string) error {
    var err error
    Pool, err = pgxpool.New(ctx, uri)
    if err != nil {
        log.Fatal("Unable to connect to database:", err)
    }

    if err := Pool.Ping(ctx); err != nil {
        log.Fatal("Unable to ping database:", err)
    }

    fmt.Println("Connected to PostgreSQL database!")
	return err
}

func DisconnectDB() {
    if Pool != nil {
        Pool.Close()
        fmt.Println("Disconnected from PostgreSQL database.")
    }
}