package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FoodItem struct {
	ID             primitive.ObjectID `bson:"_id"`
	FoodItem_ID    string             `json:"fooditem_id"`
	FoodItem_Name  *string            `json:"fooditem_name"`
	FoodItem_Price *float64           `json:"fooditem_price"`
	FoodITemImage  *string            `json:"fooditem_image"`
	Menu_Id        *string            `json:"menu_id"`
	Created_at     time.Time          `json:"created_at"`
	Updated_at     time.Time          `json:"updated_at"`
}
