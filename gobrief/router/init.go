package router

import (
	"github.com/gin-gonic/gin"
	"gobrief/gobrief/middleware"
	v1 "gobrief/gobrief/router/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	// 不涉及权限/认证 中间件
	middleware.InitMiddleware(r)
	v1.InitV1Router(r)
	return r
}