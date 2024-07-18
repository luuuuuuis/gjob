package service

import "github.com/gin-gonic/gin"

func Login(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"errcode": 0,
		"data":    "请求成功",
	})
}
