package database

import (
	"Uwall/config"
	"Uwall/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Init() {
	dsn := config.Username + ":" + config.Password + "@tcp(" + config.Address + ":" + config.Port + ")/" + config.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	Db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	Db.AutoMigrate(&model.Message{})
}
