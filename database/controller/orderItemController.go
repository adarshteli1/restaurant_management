package controller

import (
	"restaurant_management/database"

	"go.mongodb.org/mongo-driver/mongo"
)

var orderItemCollection *mongo.Collection = database.OpenCollection(database.Client, "orderitem")
