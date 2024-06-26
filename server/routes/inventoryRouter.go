package routes

import (
	controller "stock-management-server/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/items", controller.GetUsers())
	incomingRoutes.GET("/items/:item_name", controller.GetUser())
	incomingRoutes.GET("/items/:item_id", controller.GetUser())

}
