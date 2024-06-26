package routes

import (
	controller "stock-management-server/controllers"

	"github.com/gin-gonic/gin"
)

func ReturnItemsRouter(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/lend_items", controller.GetUsers())
	incomingRoutes.GET("lend_items/:items_id", controller.GetUsers())

}
