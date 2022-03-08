package routers

import (
	"feed/meta/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/api/get", controller.Get)
	return router
}
