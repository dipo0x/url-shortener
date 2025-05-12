package models

import (
	"time"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type URL struct {
	ID        uuid.UUID `bson:"_id,omitempty" json:"id"`
	URL       string    `bson:"url" json:"url" validate:"required"`
	ShortURL  string     `bson:"shortURL" json:"shortURL"`
	Clicks    int64     `bson:"clicks" json:"clicks"`
	ExpiresAt primitive.DateTime `bson:"expiresAt" json:"expiresAt" validate:"required"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}
