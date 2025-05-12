package controller

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/dipo0x/golang-url-shortener/api"
	"github.com/dipo0x/golang-url-shortener/helpers"
	"github.com/dipo0x/golang-url-shortener/internal/config"
	"github.com/dipo0x/golang-url-shortener/internal/models"
	"github.com/dipo0x/golang-url-shortener/internal/types"
	"github.com/dipo0x/golang-url-shortener/jobs"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/hibiken/asynq"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateURL(c *fiber.Ctx) error {
	var body types.IURL
	var url models.URL

	if err := c.BodyParser(&body); err != nil {
		return api.RespondWithError(c, fiber.StatusBadRequest, "Invalid request payload")
	}

	if body.URL == "" || body.ExpiresAt == "" {
		return api.RespondWithError(c, fiber.StatusBadRequest, "URL and expiresAt fields are required")
	}

	shortURL, err := helpers.GetUniqueRandomString()
	if err != nil {
		return api.RespondWithError(c, fiber.StatusInternalServerError, "Failed to generate short URL")
	}

	filter := bson.M{"url": body.URL}
	err = config.MongoDatabase.Collection("urls").FindOne(context.TODO(), filter).Decode(&url)

	if err == nil {
		shortURL := fmt.Sprintf("%s/%s", config.Config("URL_ABSOLUTE_URL"), url.ShortURL)

		return api.RespondWithError(c, fiber.StatusBadRequest, fmt.Sprintf("URL has already been shortened: %s", shortURL))
	} else if err != mongo.ErrNoDocuments {
		return api.RespondWithError(c, fiber.StatusInternalServerError, "Database error")
	}

	hours, err := time.ParseDuration(body.ExpiresAt + "h")
	if err != nil {
		return api.RespondWithError(c, fiber.StatusBadRequest, "Invalid expiration format")
	}

	urlId := uuid.New()

	urlData := models.URL{
		ID:        urlId,
		ShortURL:  shortURL,
		URL:       body.URL,
		ExpiresAt: primitive.NewDateTimeFromTime(time.Now().Add(hours)),
		Clicks:    0,
	}

	_, err = config.MongoDatabase.Collection("urls").InsertOne(context.Background(), urlData)
	if err != nil {
		return api.RespondWithError(c, fiber.StatusInternalServerError, "Failed to save URL")
	}

	hoursInt, err := strconv.Atoi(body.ExpiresAt)
	if err != nil {
		log.Fatalf("Invalid ExpiresAt value: %v", err)
	}

	delay := time.Duration(hoursInt) * time.Hour

	asynqClient := config.AsynqClient

	task, err := jobs.NewDeleteTask(urlId.String())
	if err != nil {
		return fiber.ErrInternalServerError
	}
	if _, err := asynqClient.Enqueue(task, asynq.ProcessIn(delay)); err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"success": true,
		"data":    urlData,
	})
}

func RedirectURL(c *fiber.Ctx) error {
	var url models.URL
	var err error

	id := c.Params("id")

	filter := bson.M{"shortURL": id}
	err = config.MongoDatabase.Collection("urls").FindOne(context.TODO(), filter).Decode(&url)

	if err == mongo.ErrNoDocuments {
		return api.RespondWithError(c, fiber.StatusNotFound, "Short URL not found")
	} else if err != nil {
		return api.RespondWithError(c, fiber.StatusInternalServerError, "Database error")
	}

	return c.Redirect(url.URL, fiber.StatusMovedPermanently)
}
