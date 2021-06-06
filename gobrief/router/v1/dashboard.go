package v1

import (
	"github.com/gin-gonic/gin"
	"gobrief/app/api"
)

func InitDashboardRouter(g *gin.RouterGroup)  {
	rt := g.Group("/dashboard")
	{
		rt.GET("", api.Dashboard)
	}
}