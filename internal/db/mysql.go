package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitMysqlDB() error {
	dsn := "root:2048711712P!@tcp(127.0.0.1:3306)/review_system?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database:%v", err)
		return err
	}
	return nil
}
