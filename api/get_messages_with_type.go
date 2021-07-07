package api

import (
	"Uwall/database"
	"Uwall/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetMessagesWithType(c *gin.Context) {
	messageType := c.Param("mtype")
	fmt.Println(messageType)
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))
	var messages []model.Message
	if result := database.Db.Model(&model.Message{}).Preload("Comments").Where("status = ? AND type = ?", "1", messageType).Limit(limit).Offset(offset).Find(&messages); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusOK, messages)
}
