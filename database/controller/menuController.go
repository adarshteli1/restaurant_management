package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"restaurant_management/database"
	"restaurant_management/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")

func GetMenuId() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		result, err := menuCollection.Find(context.TODO(), bson.M{})

		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro while listinhg the menu"})
			return
		}
		var allmenus []bson.M

		if err = result.All(ctx, &allmenus); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allmenus)
	}
}
func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		menu_Id := c.Param("menu_id")

		var menu models.Menu
		err := foodItemCollection.FindOne(ctx, bson.M{"menu_id": menu_Id}).Decode(&menu)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "FoodItem not found"})
		}
		c.JSON(http.StatusOK, menu)
	}
}
func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var menu models.Menu
		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error binding with JSON"})
		}

		validateErr := validate.Struct(menu)
		if validateErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validateErr.Error()})
		}

		menu.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.ID = primitive.NewObjectID()
		menu.Menu_ID = menu.ID.Hex()

		result, inserErr := menuCollection.InsertOne(ctx, menu)
		if inserErr != nil {
			msg := fmt.Sprintf("FoodItem Not Inserted")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		defer cancel()
		c.JSON(http.StatusOK, result)
		defer cancel()

	}
}

func inTimeSpan(start, end, check time.Time) bool {
	return start.After(time.Now()) && end.After(start)
}
func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var menu models.Menu
		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error binding with JSON"})
		}

		menuId := c.Param("menu_id")
		filter := bson.M{"menu_id": menuId}

		var UpdateObj primitive.D

		if menu.Start_Date != nil && menu.End_Date != nil {
			if !inTimeSpan(*menu.Start_Date, *menu.End_Date, time.Now()) {
				msg := "kindly retype time"
				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
				defer cancel()
				return
			}

			UpdateObj = append(UpdateObj, bson.E{Key: "start_date", Value: menu.Start_Date})
			UpdateObj = append(UpdateObj, bson.E{Key: "end_date", Value: menu.End_Date})

			if menu.Menu_Name != "" {
				UpdateObj = append(UpdateObj, bson.E{Key: "menu_name", Value: menu.Menu_Name})
			}

			menu.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC1123))
			UpdateObj = append(UpdateObj, bson.E{"updated_at", menu.Updated_at})

			upsert := true
			opt := options.UpdateOptions{
				Upsert: &upsert,
			}
			result, err := menuCollection.UpdateOne(
				ctx,
				filter,
				bson.D{
					{Key: "$set", Value: UpdateObj},
				},
				&opt,
			)
			if err != nil {
				msg := "Menu Not Updated"
				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
				return
			}
			defer cancel()
			c.JSON(http.StatusOK, result)

		}

	}
}
