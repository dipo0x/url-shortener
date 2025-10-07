package controller

import (
	"context"
	"fmt"
	"log"

	"time"

	"github.com/dipo0x/golang-url-shortener/api"
	"github.com/dipo0x/golang-url-shortener/helpers"
	"github.com/dipo0x/golang-url-shortener/internal/config"
	"github.com/dipo0x/golang-url-shortener/internal/models"
	"github.com/dipo0x/golang-url-shortener/internal/types"
	queue "github.com/dipo0x/golang-url-shortener/workers/rabbitmq"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

var ctx = context.Background()

func CreateURL(c *fiber.Ctx) error {
	var body types.IURL
	// var url models.URL

	if err := c.BodyParser(&body); err != nil {
		return api.RespondWithError(c, fiber.StatusBadRequest, "Invalid request payload")
	}

	if body.URL == "" || body.ExpiresAt == "" {
		return api.RespondWithError(c, fiber.StatusBadRequest, "URL and expiresAt fields are required")
	}

	shortURL, err := helpers.GetUniqueRandomString()
	if err != nil {
		log.Fatal(err)
		return api.RespondWithError(c, fiber.StatusInternalServerError, "Failed to generate short URL")
	}

	var url_short_form string
	err = config.Pool.QueryRow(ctx, `
		SELECT short_url FROM urls WHERE url = $1 LIMIT 1
	`, body.URL).Scan(&url_short_form)

	if err == nil {
		shortURL := fmt.Sprintf("%s/%s", config.Config("URL_ABSOLUTE_URL"), url_short_form)

		return api.RespondWithError(c, fiber.StatusBadRequest, fmt.Sprintf("URL has already been shortened: %s", shortURL))
	}

	hours, err := time.ParseDuration(body.ExpiresAt + "h")
	if err != nil {
		return api.RespondWithError(c, fiber.StatusBadRequest, "Invalid expiration format")
	}

	urlId := uuid.New()

	var savedURL models.URL

	err = config.Pool.QueryRow(ctx, `
	INSERT INTO urls (id, url, short_url, clicks, expires_at)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, url, short_url, clicks, expires_at, created_at, updated_at
`, urlId, body.URL, shortURL, 0, time.Now().Add(hours)).Scan(
    &savedURL.ID,
    &savedURL.URL,
    &savedURL.ShortURL,
    &savedURL.Clicks,
    &savedURL.ExpiresAt,
    &savedURL.CreatedAt,
    &savedURL.UpdatedAt,
)

	if err != nil {
		return api.RespondWithError(c, fiber.StatusInternalServerError, "Failed to save URL")
	}

	queue.PublishJob(urlId.String(), float64(hours))

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": 200,
		"success": true,
		"data": savedURL,
	})
}

func RedirectURL(c *fiber.Ctx) error {
	id := c.Params("id")
	var originalURL string

    err := config.Pool.QueryRow(ctx, `
        UPDATE urls
        SET clicks = clicks + 1, updated_at = NOW()
        WHERE short_url = $1
        RETURNING url
    `, id).Scan(&originalURL)

    if err != nil {
        if err == pgx.ErrNoRows {
            return api.RespondWithError(c, fiber.StatusInternalServerError, "URL not found")
        }
    }

	return c.Redirect(originalURL, fiber.StatusMovedPermanently)
}

