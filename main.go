package main

import (
	"github.com/dipo0x/golang-url-shortener/config"
	"github.com/dipo0x/golang-url-shortener/routes"
	"log"
	"github.com/gofiber/fiber/v2"
)

func main () {

	app := fiber.New()
	api := app.Group("/api")

	routes.IndexRoutes(api.Group("/index"))
	routes.URLRoutes(api.Group("/url"))

	err:= config.InitializeMongoDB(config.Config("MONGO_URI"), config.Config("MONGO_DATABASE"))

	if err != nil {
		defer config.DisconnectMongoDB()
		log.Fatalf("Could not connect to MongoDB: %v", err)
	}
	port := config.Config("PORT")
	log.Fatal(app.Listen(port))
}