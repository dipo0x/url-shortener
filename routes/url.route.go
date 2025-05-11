package routes

import (
	"github.com/gofiber/fiber/v2"
    "github.com/dipo0x/golang-url-shortener/controller"
)

func URLRoutes(router fiber.Router) {
    router.Post("/create-url", controller.CreateURL)
    router.Get("/:id", controller.RedirectURL)
}