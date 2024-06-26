package main

import (
	"os"

	middleware "stock-management-server/middleware"
	routes "stock-management-server/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.UserRoutes(router)
	routes.Use(middleware.Authentication())

	routes.InventoryRoutes(router)
	routes.RequestItemsRoutes(router)
	routes.ReturnItemsRoutes(router)

	router.Run(":" + port)

}
