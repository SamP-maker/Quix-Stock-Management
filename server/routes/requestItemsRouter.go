package routes

import (
	controller "stock-management-server/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/request_form", controller.GetUsers())
	incomingRoutes.GET("/request_form/:form_id", controller.GetUser())

	incomingRoutes.POST("/users/form", controller.Login())
}
