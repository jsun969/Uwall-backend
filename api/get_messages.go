package api

import (
	"Uwall/database"
	"Uwall/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetMessages(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))
	var messages []model.ResMessage
	if result := database.Db.Model(&model.Message{}).Where("status = ?", "1").Limit(limit).Offset(offset).Find(&messages); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusOK, messages)
}
