package web

import (
	"fmt"
	"general-service/library/resource"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type LoginInfo struct {
	// binding中的即为条件,当需要同时满足的时候,用","来进行间隔;当是或的关系的时候,使用"|"
	UserName string `form:"username" binding:"required"`
	PassWord string `form:"password" binding:"required"`
}

// 登陆接口
func LoginEndpoint(ctx *gin.Context) {

	var loginInfo LoginInfo
	//校验参数的先决条件在于先要对参数与结构体进行绑定
	err := ctx.ShouldBind(&loginInfo)

	if err != nil {
		ctx.String(http.StatusBadRequest, "the param not valid, err: %V", err)
		//如果下面不加return的话,函数仍然会继续执行
		return
	}

	redisErr := resource.RedisClient.Set(ctx, loginInfo.UserName, loginInfo.PassWord,
		time.Duration(60)*time.Second).Err()
	if redisErr != nil {
		log.Fatalf("redis操作失败%s", redisErr.Error())
	}

	log.Printf("username:%s, password:%s", loginInfo.UserName, loginInfo.PassWord)

	ctx.String(200, fmt.Sprintf("user:%s, welcome", loginInfo.UserName))
}
