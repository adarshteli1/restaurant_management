package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderModel struct {
	ID         primitive.ObjectID `json:"_id"`
	Order_Id   string             `json:"order_id"`
	Table_Id   *string            `json:"table_id"`
	Order_Date time.Time          `json:"order_date"`
	Creted_at  time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
}
