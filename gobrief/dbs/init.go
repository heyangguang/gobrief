package dbs

import (
	"gobrief/gobrief/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"time"
)

var Orm *gorm.DB

func InitDB()  {
	Orm = gormDB()
}

func gormDB() *gorm.DB {
	dsn := "root:123456@tcp(localhost:3306)/go_brief?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("连接数据库失败:"+ err.Error())
		os.Exit(-1)
	}
	mysqlDB, err := db.DB()
	if err != nil {
		logger.Error("返回数据库对象失败:"+ err.Error())
		os.Exit(-1)
	}
	mysqlDB.SetMaxIdleConns(5)
	mysqlDB.SetMaxOpenConns(10)
	mysqlDB.SetConnMaxLifetime(time.Second*30)
	return db
}
