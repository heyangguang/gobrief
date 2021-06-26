package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"gobrief/gobrief/dbs"
	. "gobrief/gobrief/form_validation"
	"gobrief/gobrief/logger"
	"gobrief/gobrief/router"
)

var configPath string

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start go-brief port default 8080",
	Run: func(cmd *cobra.Command, args []string) {
		// 初始化配置文件
		viperInit(configPath)
		fmt.Println(AllConfig)
		runStart()
	},
}

func runStart() {
	// 日志
	err := logger.InitLogger(AllConfig.LogPath, 2048, 30, 200, AllConfig.LogLevel)
	if err != nil {
		panic(err.Error())
	}

	// 数据库
	dbs.InitDB(AllConfig.DBConnect)
	dbs.InitTable()

	// 表单验证
	InitFormValidation()

	// !DEBUG
	if AllConfig.LogLevel != "DEBUG" {
		gin.SetMode(gin.ReleaseMode)
		logger.Info(fmt.Sprintf("Listening and serving HTTP on :%d", AllConfig.Port))
	}

	r := router.InitRouter()
	// TODO:后面会增加容错与优雅退出
	_ = r.Run(fmt.Sprintf(":%d", AllConfig.Port))
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringVar(&configPath, "config", "./config_dev.yaml", "--config ./config_prod.yaml")
}
