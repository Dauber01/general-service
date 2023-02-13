// 启动程序文件
package main

import (
	"context"
	"general-service/httpserver"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"general-service/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 初始化 server信息
	s := httpserver.InitServer(router)

	// 设置路由信息
	httpserver.ReqHandler(router)

	// 初始化各个中间件模版
	ctx := context.Background()
	bootstrap.ResourceInit(ctx)

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 3 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
