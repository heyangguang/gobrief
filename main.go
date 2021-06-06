package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gobrief/gobrief/dbs"
	. "gobrief/gobrief/form_validation"
	"gobrief/gobrief/logger"
	"gobrief/gobrief/router"
)

func main()  {
	// 日志
	logLevel := "DEBUG"
	port := 8080
	err := logger.InitLogger("logs/go-brief.log", 2048, 30, 200, logLevel)
	if err != nil {
		panic(err.Error())
	}

	// 数据库
	dbs.InitDB()
	dbs.InitTable()

	// 表单验证
	InitFormValidation()

	// !DEBUG
	if logLevel != "DEBUG" {
		gin.SetMode(gin.ReleaseMode)
		logger.Info(fmt.Sprintf("Listening and serving HTTP on :%d", port))
	}

	r := router.InitRouter()
	r.Run(":8080")
}