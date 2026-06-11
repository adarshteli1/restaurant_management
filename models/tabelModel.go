package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Table struct {
	ID           primitive.ObjectID `json:"_id"`
	Table_Id     string             `json:"table_id"`
	Table_Number *int               `json:"table_number"`
	Table_Guest  *int               `json:"table_guest"`
	Creted_at    time.Time          `json:"created_at"`
	Updated_at   time.Time          `json:"updated_at"`
}
