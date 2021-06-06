package middleware

import (
	"github.com/gin-gonic/gin"
	"gobrief/gobrief/logger"
)

func InitMiddleware(r *gin.Engine) {

	r.Use(GinLogger(logger.Logger))

	r.Use(GinRecovery(logger.Logger, true))

}
