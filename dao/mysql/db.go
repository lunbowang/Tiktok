package mysql

import (
	"fmt"
	"log"

	"gorm.io/gorm/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Init() {
	// 配置 MySQL
	username := "root"
	password := "20050606a"
	host := "127.0.0.1"
	port := 3306
	dbname := "tiktok"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln("数据库连接失败", err)
	}
	Db = db
}
