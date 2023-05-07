package httpserver

import (
	"general-service/library/resource"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 初始化 server
func InitServer(router *gin.Engine) *http.Server {
	return &http.Server{
		Addr:           resource.Conf.Server.Addr,
		Handler:        router,
		ReadTimeout:    time.Duration(resource.Conf.Server.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(resource.Conf.Server.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
