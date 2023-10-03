package routers

import (
	"github.com/gin-gonic/gin"
	orderController "github.com/mcsans/assignment2-019/controllers"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	
	router.GET("/orders", orderController.Index)
	router.GET("/orders/:id", orderController.Show)
	router.POST("/orders", orderController.Create)
	router.PUT("/orders/:id", orderController.Update)
	router.DELETE("/orders/:id", orderController.Delete)

	return router
}