package main

import (
	"Uwall/api"
	"Uwall/config"
	"Uwall/model/database"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := config.Username + ":" + config.Password + "@tcp(" + config.Address + ":" + config.Port + ")/" + config.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&database.Message{})

	r := gin.Default()
	r.POST("/love", api.PostLove)
	r.Run()
}
