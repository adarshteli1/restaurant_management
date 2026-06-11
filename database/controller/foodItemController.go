package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"restaurant_management/database"
	"restaurant_management/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodItemCollection *mongo.Collection = database.OpenCollection(database.Client, "fooditem")
var validate = validator.New()

func GetFoodItemId() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		foodItemId := c.Param("fooditem_id")

		var fooditem models.FoodItem
		err := foodItemCollection.FindOne(ctx, bson.M{"fooditem_id": foodItemId}).Decode(&fooditem)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "FoodItem not found"})
		}
		c.JSON(http.StatusOK, fooditem)
	}
}

func CreateFoodItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var menu models.Menu
		var foodItem models.FoodItem

		if err := c.BindJSON(&foodItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error binding with JSON"})
		}

		validateErr := validate.Struct(foodItem)
		if validateErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validateErr.Error()})
		}

		err := menuCollection.FindOne(ctx, bson.M{"menu_id": foodItem.Menu_Id}).Decode(&menu)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "FoodITem Not Found	"})
		}

		foodItem.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		foodItem.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		foodItem.ID = primitive.NewObjectID()
		foodItem.FoodItem_ID = foodItem.ID.Hex()
		var num = toFixed(*&foodItem.FoodItem_Price, 2)
		foodItem.FoodItem_Price = &num

		result, inserErr := foodItemCollection.InsertOne(ctx, foodItem)
		if inserErr != nil {
			msg := fmt.Sprintf("FoodItem Not Inserted")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

func GetFoodItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
		if err != nil || recordPerPage < 1 {
			recordPerPage = 10
		}
		page, err1 := strconv.Atoi(c.Query("page"))
		if err1 != nil || page < 1 {
			page = 1
		}
		startIndex := (page - 1) * recordPerPage
		startIndex, err = strconv.Atoi(c.Query("startIndex"))
		matchStage := bson.D{
			{Key: "$match", Value: bson.D{}},
		}

		groupStage := bson.D{
			{
				Key: "$group",Value: bson.D{
					{
						Key: "_id",Value: bson.D{
							{Key:"_id",Value: "null"}}},
							{Key: "total_count",Value: bson.D{{"$sum,1"}
						}
					},
					{Key: "data",Value: bson.D{{Key: "$push",Value: "$$ROOT"}}}
				}
			}
		}
		projectStage := bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "_id", Value: 0},
				{Key: "total_count", Value: 1},
				{Key: "food_items", Value: bson.D{
					{Key: "$slice", Value: []interface{}{"$data", startIndex, recordPerPage}},
				}},
			}},
		}
		result, err := foodItemCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage,
			groupStage,
			projectStage,
		})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}

		var allFoodItems []bson.M
		if err=result.All(ctx,&allFoodItems);err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, allFoodItems[0])

	}

}

func UpdateFoodItem() gin.HandlerFunc {
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var menu models.Menu
		var foodItem models.FoodItem

		if err := c.BindJSON(&foodItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error binding with JSON"})
		}


		var UpdateObj primitive.D
	}
}
