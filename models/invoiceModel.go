package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InvoiceModel struct {
	ID               primitive.ObjectID `bson:"_id"`
	Invoice_Id       string             `json:"invoice_id"`
	Order_Id         string             `json:"order_Id"`
	Payment_Method   *string            `json:"payment_method"`
	Payment_Status   *string            `json:"payment_status"`
	Payment_Due_Date string             `json:"payment_due_date"`
	Creted_at        time.Time          `json:"created_at"`
	Updated_at       time.Time          `json:"updated_at"`
}
