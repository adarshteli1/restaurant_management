package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderItem struct {
	ID                   primitive.ObjectID `json:"_id"`
	Order_Id             string             `json:"order_id"`
	FoodItem_ID          *string            `json:"fooditem_id"`
	OrderItem_Id         string             `json:"orderitem_id"`
	OrderItem_Quantity   *string            `json:"orderitem_quantity"`
	OrderItem_Unit_Price *float64           `json:"orderitem_unit_price"`
	Creted_at            time.Time          `json:"created_at"`
	Updated_at           time.Time          `json:"updated_at"`
}
