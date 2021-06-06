package v1

import (
	"github.com/gin-gonic/gin"
	"gobrief/app/api"
)

func InitBaseRouter(g *gin.RouterGroup)  {
	rt := g.Group("")
	{
		rt.GET("/login", api.Dashboard)
	}
}
