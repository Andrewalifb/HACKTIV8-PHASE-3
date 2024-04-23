package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Employee struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string             `json:"firstname"`
	LastName  string             `json:"lastname"`
	Email     string             `json:"email"`
	Phone     string             `json:"phone"`
	Position  string             `json:"position"`
	CreatedAt time.Time          `json:"createdat"`
	UpdatedAt time.Time          `json:"updatedat"`
}

