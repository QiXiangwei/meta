package service

import (
	"feed/meta/library/logs"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Api struct {
	
}

func (a *Api) Get(ctx *gin.Context) {
	id := ctx.DefaultQuery("id", "0")
	if id == "0" {
		logs.Logger().WithFields(logrus.Fields{
			"errno": 100,
		}).Warn("id is empty!")
	}
}
