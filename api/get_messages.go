package api

import (
	"Uwall/database"
	"Uwall/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetMessages(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	var messages []model.Message
	if result := database.Db.Model(&model.Message{}).Preload("Comments", func(db *gorm.DB) *gorm.DB {
		return db.Where("status = ?", 1).Order("created_at desc")
	}).Where("status = ?", "1").Order("created_at desc").Limit(10).Offset(page * 10).Find(&messages); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusOK, messages)
}
