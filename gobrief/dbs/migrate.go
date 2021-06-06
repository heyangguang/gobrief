package dbs

import (
	"gobrief/app/model"
	"gobrief/gobrief/logger"
	"os"
)

func InitTable() {
	err := Orm.AutoMigrate(&model.UserModel{})
	if err != nil {
		logger.Error("初始化数据库表失败:" + err.Error())
		os.Exit(-1)
	}
}
