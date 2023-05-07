// 启动程序文件
package main

import (
	"context"
	"general-service/httpserver"
	"general-service/library/resource"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"general-service/bootstrap"

	"general-service/httpserver/pkg/middleware"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {

	// 初始化配置信息
	initConfig(context.Background())

	router := gin.Default()

	// 初始化 server信息
	s := httpserver.InitServer(router)

	// 初始化权限系统
	authMiddleware, err := middleware.AuthMiddlewareInit()
	if err != nil {
		log.Fatal("middleware.AuthMiddlewareInit() Error:" + err.Error())
	}
	errInit := authMiddleware.MiddlewareInit()
	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	// 设置安全相关的路由
	router.POST("/v1/login", authMiddleware.LoginHandler)
	router.GET("/v1/refresh_token", authMiddleware.RefreshHandler)
	router.Use(authMiddleware.MiddlewareFunc())
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

// 初始化配置文件
// https://github.com/spf13/viper
func initConfig(_ context.Context) {

	// 1.指定对应的文件
	viper.AddConfigPath("./conf/")
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	// 2.读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("read config file failed, %v", err)
	}
	var conf resource.Config
	if err := viper.Unmarshal(&conf); err != nil {
		log.Fatalf("unmarshal config file failed, %v", err)
	}
	// 3.赋值给对应的变量
	resource.Conf = &conf
	log.Printf("config load, conf:%v", resource.Conf)
	log.Printf("config load, server:%v", resource.Conf.Server)
	log.Printf("config load, redis:%v", resource.Conf.Redis)
}
