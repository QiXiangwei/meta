package controller

import (
	"feed/meta/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Get(c *gin.Context) {
	var api service.Api
	api.Get(c)
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
