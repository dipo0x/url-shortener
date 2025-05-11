package main

import (
	"log"

	"github.com/dipo0x/golang-url-shortener/app"
	"github.com/dipo0x/golang-url-shortener/config"

)

func main() {
	app := app.InitializeApp()
	port := config.Config("PORT")
	log.Fatal(app.Listen(port))
}
