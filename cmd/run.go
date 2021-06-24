package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gobrief/gobrief/dbs"
	. "gobrief/gobrief/form_validation"
	"gobrief/gobrief/logger"
	"gobrief/gobrief/router"
)

var port int
var log string
var logLevel string
var db string
var configPath string
var server bool

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start go-brief port default 8080",
	Run: func(cmd *cobra.Command, args []string) {
		// 初始化配置文件
		viperInit(configPath)
		if server {
			runStart(viper.GetString("log_path"),
				viper.GetString("log_level"),
				viper.GetString("db_connect"),
				viper.GetInt("port"))
			return
		} else {
			runStart(log, logLevel, db, port)
		}
	},
}

func runStart(logPath, logLevel, db string, port int) {
	// 日志
	err := logger.InitLogger(logPath, 2048, 30, 200, logLevel)
	if err != nil {
		panic(err.Error())
	}

	// 数据库
	dbs.InitDB(db)
	dbs.InitTable()

	// 表单验证
	InitFormValidation()

	// !DEBUG
	if logLevel != "DEBUG" {
		gin.SetMode(gin.ReleaseMode)
		logger.Info(fmt.Sprintf("Listening and serving HTTP on :%d", port))
	}

	r := router.InitRouter()
	// TODO:后面会增加容错与优雅退出
	_ = r.Run(fmt.Sprintf(":%d", port))
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringVarP(&log, "log", "l", "./logs/go-brief.log", "--log <logLocation> example --log ./logs/go-brief.log")
	runCmd.Flags().IntVarP(&port, "port", "p", 8080, "--port <port> example --port 8080")
	runCmd.Flags().StringVar(&logLevel, "logLevel", "INFO", "--logLevel <logLocation> example --logLevel INFO")
	runCmd.Flags().StringVar(&db, "db", "", "--db <logLocation> example --db root:123456@(127.0.0.1:3306)/test")
	runCmd.Flags().BoolVar(&server, "server", true, "--server")
	runCmd.Flags().StringVar(&configPath, "config", "./config_dev.yaml", "--config ./config_prod.yaml")
}
