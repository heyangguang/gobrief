package v1

import (
	"github.com/gin-gonic/gin"
)

func InitV1Router(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("")

	// 不需要认证
	PublicRouter(g)

	// 需要认证
	PrivateRouter(g)

	return g
}

func PublicRouter(g *gin.RouterGroup)  {
	publicGroup := g.Group("/api/v1")
	{
		InitBaseRouter(publicGroup)
	}
}

func PrivateRouter(g *gin.RouterGroup)  {
	privateGroup := g.Group("/api/v1")
	//privateGroup.Use(middleware.Logger2())
	{
		InitDashboardRouter(privateGroup)
		InitUserRouter(privateGroup)
	}
}
