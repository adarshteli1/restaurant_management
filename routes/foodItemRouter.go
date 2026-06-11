package routes

import (
	"restaurant_management/controller"

	"github.com/gin-gonic/gin"
)

func FoodItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/fooditem", controller.GetFoodItem())
	incomingRoutes.GET("/fooditem/:food_id", controller.GetFoodItemId())
	incomingRoutes.POST("/fooditem", controller.CreateFoodItem())
	incomingRoutes.PATCH("/fooditem", controller.UpdateFoodItem())

}
