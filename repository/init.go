package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dsn = "root:root123456@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
}
