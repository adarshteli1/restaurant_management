package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Notes struct {
	ID         primitive.ObjectID `bson:"_id"`
	Notes_Id   string             `json:"notes_id"`
	Text       string             `json:"text"`
	Title      string             `json:"title"`
	Creted_at  time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
}
