package v1

import (
	"github.com/gin-gonic/gin"
	"gobrief/app/api"
)

func InitUserRouter(g *gin.RouterGroup)  {
	rt := g.Group("/user")
	{
		rt.POST("", api.UserCreate)
	}
}
