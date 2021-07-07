package api

import (
	"Uwall/database"
	"Uwall/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostNotice(c *gin.Context) {
	var req model.ReqBaseWithName
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	message := model.Message{
		Type:     "notice",
		FromName: req.Name,
		Message:  req.Message,
		ImageUrl: req.ImageUrl,
		Status:   false,
	}
	if result := database.Db.Create(&message); result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": result.Error})
		return
	}
	c.Status(http.StatusCreated)
}
