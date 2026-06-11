package routes

import "github.com/gin-gonic/gin"

func InvoiceRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/invoice", controller.GetInvoice())
	incomingRoutes.GET("/invoice/:invoice_id", controller.GetInvoiceId())
	incomingRoutes.POST("/invoice", controller.CreateInvoice())
	incomingRoutes.PATCH("/invoice/:invoice_id", controller.UpdateInvoice())

}
