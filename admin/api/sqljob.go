package api

import (
	"fmt"
	"gjob-admin/pkg/model"
	"gjob-admin/service"
	"net/http"

	"gjob-admin/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type sqljob struct{}

var Sqljob sqljob

// 创建sqljob
func (s *sqljob) create(ctx *gin.Context) {
	var sqljobCreate = &model.SqlJob{}
	fmt.Println(sqljobCreate.WechatGroup)
	path := ctx.Request.URL.Path
	query := ctx.Request.URL.RawQuery

	if err := ctx.ShouldBindJSON(sqljobCreate); err != nil {
		utils.Logger.Error().
			Err(errors.New("参数解析失败")).
			Stack().
			Int("status", 500).
			Str("method", ctx.Request.Method).
			Str("path", path).
			Str("query", query).
			Str("ip", ctx.ClientIP()).
			Str("user-agent", ctx.Request.UserAgent()).
			Msg("请求响应超时")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	fmt.Println(sqljobCreate.WechatGroup)

	if err := service.SqljobService.Create(sqljobCreate); err != nil {
		utils.Logger.Error().
			Err(err).
			Stack().
			Int("status", 500).
			Str("method", ctx.Request.Method).
			Str("path", path).
			Str("query", query).
			Str("ip", ctx.ClientIP()).
			Str("user-agent", ctx.Request.UserAgent()).
			Msg("创建sqljob失败")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"msg":  "创建sqljob成功",
			"data": nil,
		},
	)

}

// 查询sqljob
func (s *sqljob) query(ctx *gin.Context) {
	name := ctx.Query("name")
	sqljoblist, err := service.SqljobService.Get(name)
	if err != nil {
		utils.Logger.Error().
			Err(err).
			Stack().
			Int("status", 500).
			Str("method", ctx.Request.Method).
			Str("path", ctx.Request.URL.Path).
			Str("query", ctx.Request.URL.RawQuery).
			Msg("查询sqljob失败")
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"msg":  "查询sqljob成功",
			"data": sqljoblist,
		},
	)
}

// 分页查询
func (s *sqljob) queryByPage(ctx *gin.Context) {
	pagesize := ctx.Query("pageSize")
	pagenum := ctx.Query("pageNum")
	sqljoblist, err := service.SqljobService.GetByPage(pagenum, pagesize)
	if err != nil {
		utils.Logger.Error().
			Err(err).
			Stack().
			Int("status", 500).
			Str("method", ctx.Request.Method).
			Str("path", ctx.Request.URL.Path).
			Str("query", ctx.Request.URL.RawQuery).
			Msg("分页查询sqljob失败")
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"msg":  "分页查询sqljob成功",
			"data": sqljoblist,
		},
	)
}
