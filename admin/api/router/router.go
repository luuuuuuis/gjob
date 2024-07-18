package router

import (
	"gjob-admin/api/service"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	// 创建登录路由
	r.POST("/login", service.Login)

}
