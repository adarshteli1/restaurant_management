package routes

import "github.com/gin-gonic/gin"

func TableRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/tabel", controller.GetTable())
	incomingRoutes.GET("/tabel/:table_id", controller.GetTableId())
	incomingRoutes.POST("/tabel", controller.CreateTable())
	incomingRoutes.PATCH("/tabel:table_id", controller.UpdateTable())

}
