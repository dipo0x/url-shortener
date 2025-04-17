package controller

import (
	"context"
	"time"
	"github.com/dipo0x/golang-url-shortener/config"
	"github.com/dipo0x/golang-url-shortener/helpers"
	"github.com/dipo0x/golang-url-shortener/models"
	"github.com/dipo0x/golang-url-shortener/types"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"

)

func CreateURL(c *fiber.Ctx) error {
	var body types.IURL
	var url models.URL

	if err := c.BodyParser(&body); err != nil {
		return helpers.RespondWithError(c, fiber.StatusBadRequest, "Invalid request payload")
	}
	
	if body.URL == "" || body.ExpiresAt == "" {
		return helpers.RespondWithError(c, fiber.StatusBadRequest, "URL and expiresAt fields are required")
	}
	shortURL, err := helpers.GetUniqueRandomString()

	filter := bson.M{"URL": body.URL}
	err = config.MongoDatabase.Collection("urls").FindOne(context.Background(), filter).Decode(&url)
	println(err)
	println(filter)
	hours, err := time.ParseDuration(body.ExpiresAt + "h")
	if err == nil {
		return helpers.RespondWithError(c, fiber.StatusBadRequest, "URL has already been shortened")
	} else if err != mongo.ErrNoDocuments {
		return helpers.RespondWithError(c, fiber.StatusInternalServerError, "Database error")
	}
	expirationTime := time.Now().Add(hours)

	urlData := models.URL{
		ID : uuid.New(),
		ShortURL:  shortURL,
		URL: body.URL,
		ExpiresAt: primitive.NewDateTimeFromTime(expirationTime),
	}

	_, err = config.MongoDatabase.Collection("urls").InsertOne(context.Background(), urlData)

	
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "status": 200,
        "success": true,
        "data": urlData,
    })
}