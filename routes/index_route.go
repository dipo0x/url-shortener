package routes

import (
	"github.com/gofiber/fiber/v2"
    "github.com/dipo0x/golang-url-shortener/controller"
)

func IndexRoutes(router fiber.Router) {
    router.Get("/", controller.Index)
}