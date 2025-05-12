package test

import (
    "log"
    "os"
    "testing"
    "github.com/dipo0x/golang-url-shortener/app"
    "github.com/joho/godotenv"
    "github.com/gofiber/fiber/v2"
)

var testApp *fiber.App

func TestMain(m *testing.M) {

    err := godotenv.Load("../.env")
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    testApp = app.InitializeApp()

    code := m.Run()

    os.Exit(code)
}