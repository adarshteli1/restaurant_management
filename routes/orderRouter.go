package routes

import "github.com/gin-gonic/gin"

func OrderRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/order", controller.GetOrder())
	incomingRoutes.GET("/order/:order_id", controller.GetOrderId())
	incomingRoutes.POST("/order", controller.CreateOrder())
	incomingRoutes.PATCH("/order", controller.UpdateOrder())
	incomingRoutes.DELETE("/order/:order_id", controller.DeleteOrder())

}
