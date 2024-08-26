package api

import (
	"gjob-admin/service"
	"net/http"

	"fmt"

	"github.com/gin-gonic/gin"
)

var User user

type user struct{}

// 验证账号密码
func (l *user) auth(ctx *gin.Context) {
	params := new(struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	})
	if err := ctx.ShouldBindJSON(params); err != nil {
		fmt.Println("Bind请求参数失败")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	err := service.Login.Auth(params.UserName, params.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "登录成功",
		"data": nil,
	})
}
