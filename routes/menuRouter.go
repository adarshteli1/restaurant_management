package routes

import (
	"restaurant_management/controller"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/menu", controller.GetMenu())
	incomingRoutes.GET("/menu/:menu_id", controller.GetMenuId())
	incomingRoutes.POST("menu", controller.CreateMenu())
	incomingRoutes.PATCH("menu/:menu_id", controller.UpdateMenu())

}
