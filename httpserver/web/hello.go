package web

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// 登陆接口
func HelloEndpoint(ctx *gin.Context) {

	// 校验参数的先决条件在于先要对参数与结构体进行绑定
	//err := ctx.ShouldBind(&loginInfo)
	name := ctx.Param("name")

	/*if err != nil {
		ctx.String(http.StatusBadRequest, "the param not valid, err: %V", err)
		// 如果下面不加return的话,函数仍然会继续执行
		return
	}*/

	log.Printf("the user's name is :%s", name)

	ctx.String(200, fmt.Sprintf("user:%s, welcome", name))
}
