package httpserver

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 初始化 server
func InitServer(router *gin.Engine) *http.Server {
	return &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
