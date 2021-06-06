package api

import (
	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {
	c.JSON(200, gin.H{"message": "0"})
}