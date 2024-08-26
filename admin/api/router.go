package api

import (
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	// 创建登录路由
	r.GET("/hello", hello)
	r.POST("/api/login", User.auth)
	r.POST("/job/sql/create", Sqljob.create)
	r.GET("/job/sql/get", Sqljob.query)
	r.GET("/job/sql/getall", Sqljob.queryAll)
	r.GET("/job/sql/getbypage", Sqljob.queryByPage)
}
