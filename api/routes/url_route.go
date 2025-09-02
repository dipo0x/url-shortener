package routes

import (
	"github.com/gofiber/fiber/v2"
    "github.com/dipo0x/golang-url-shortener/internal/controller"
)

func URLRoutes(router fiber.Router) {
     router.Get("/:id", controller.RedirectURL)
    router.Post("/create-url", controller.CreateURL)
   
}