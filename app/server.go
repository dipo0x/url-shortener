package app

import (

	"github.com/dipo0x/golang-url-shortener/internal/config"
	"github.com/dipo0x/golang-url-shortener/internal/infra"
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

	config.InitializeRabbitMQ(config.Config("RABBIT_MQ_URL"))
	err := config.InitializeDB(config.Config("DATABASE_URL"))
	
	if err != nil {
        infra.FailOnError(err, "DB init failed")
    }

	return app
}
