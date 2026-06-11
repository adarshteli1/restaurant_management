package routes

import (
	"restaurant_management/controller"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderitem", controller.GetOrderItem())
	incomingRoutes.GET("/orderitem/:orderitem_id", controller.GetOrderItemId())
	incomingRoutes.GET("/orderitem_order/:orderitem_id", controller.GetOrderItemByOrder())
	incomingRoutes.POST("/orderitem", controller.CreateOrderItem())
	incomingRoutes.PATCH("/orderitem/:orderitem_id", controller.UpdateOrderItem())

}
