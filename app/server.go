package app

import (
	"log"
	"os/exec"
	"github.com/dipo0x/golang-url-shortener/config"
	"github.com/dipo0x/golang-url-shortener/routes"
	"github.com/gofiber/fiber/v2"
)

func InitializeApp() *fiber.App {
	app := fiber.New()
	api := app.Group("/api")

	routes.IndexRoutes(api.Group("/index"))
	routes.URLRoutes(api.Group("/url"))

	config.InitializeRedis(config.Config("REDIS_URL"))
	err := config.InitializeMongoDB(config.Config("MONGO_URI"), config.Config("MONGO_DATABASE"))
	
	if err != nil {
		defer config.DisconnectMongoDB()
		log.Fatalf("Could not connect to MongoDB: %v", err)
	}

	go func() {
		cmd := exec.Command("go", "run", "workers/redis.worker.go")
		cmd.Stdout = log.Writer()
		cmd.Stderr = log.Writer()
		if err := cmd.Start(); err != nil {
			log.Fatalf("Failed to start Redis worker: %v", err)
		}
		log.Println("Redis worker started")
	}()

	return app
}
