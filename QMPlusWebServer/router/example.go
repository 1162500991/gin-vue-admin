package router

import (
	"github.com/gin-gonic/gin"
	"main/controller/api"
)

func InitExampleRouter(Router *gin.Engine) {
	BaseRouter := Router.Group("exp")
	{
		BaseRouter.POST("createExp", api.CreateExp)
	}
}
