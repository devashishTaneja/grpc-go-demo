package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const dsn = "root:@tcp(localhost:3306)/demo?charset=utf8&parseTime=True&loc=Local"

func ConnectSql() (*gorm.DB, error) {

	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}
