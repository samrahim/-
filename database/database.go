package database

import (
	"fmt"
	"learn/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBinstance struct {
	DB *gorm.DB
}

var Database DBinstance

func GetDb() {
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/demo_gorm?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
	db.Logger = db.Logger.LogMode(logger.Info)
	db.AutoMigrate(&models.User{})
	Database = DBinstance{DB: db}
}
