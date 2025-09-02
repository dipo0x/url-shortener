package app

import (
	"log"
	"os/exec"
	"github.com/dipo0x/golang-url-shortener/internal/config"
	"github.com/dipo0x/golang-url-shortener/api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func InitializeApp() *fiber.App {
	app := fiber.New()
	app.Use(logger.New())
	api := app.Group("/api")

	routes.IndexRoutes(api.Group("/index"))
	routes.URLRoutes(api.Group("/url"))

	config.InitializeRedis(config.Config("REDIS_URL"))
	err := config.InitializeDB(config.Config("DATABASE_URL"))
	
	 if err != nil {
        log.Fatalf("DB init failed: %v", err)
    }
    // defer config.DisconnectDB()

	go func() {
		cmd := exec.Command("go", "run", "workers/redis_worker.go")
		cmd.Stdout = log.Writer()
		cmd.Stderr = log.Writer()
		if err := cmd.Start(); err != nil {
			log.Fatalf("Failed to start Redis worker: %v", err)
		}
		log.Println("Redis worker started")
	}()

	return app
}
