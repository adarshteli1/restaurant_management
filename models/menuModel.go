package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct {
	ID            primitive.ObjectID `bson:"_id"`
	Menu_ID       string             `json:"menu_id"`
	Menu_Name     string             `json:"menu_name"`
	Menu_Category string             `json:"menu_category"`
	Start_Date    *time.Time         `json:"start_date"`
	End_Date      *time.Time         `json:"end_date"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
}
