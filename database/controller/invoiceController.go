package controller

import (
	"restaurant_management/database"

	"go.mongodb.org/mongo-driver/mongo"
)

var invoiceCollection *mongo.Collection = database.OpenCollection(database.Client, "invoice")
