package httpserver

import (
	"general-service/httpserver/web"

	"github.com/gin-gonic/gin"
)

func ReqHandler(router *gin.Engine) {
	// 简单的路由组: v1
	v1 := router.Group("/v1")
	{
		v1.GET("/hello", web.HelloEndpoint)
	}
}
