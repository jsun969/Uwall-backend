package api

import (
	"Uwall/database"
	"Uwall/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostHelp(c *gin.Context) {
	var req model.Base
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	message := model.Message{
		Type:      "help",
		FromName:  req.Name,
		Message:   req.Message,
		Anonymous: req.Anonymous,
		ImageUrl:  req.ImageUrl,
		Status:    false,
	}
	if result := database.Db.Create(&message); result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": result.Error})
		return
	}
	c.Status(http.StatusCreated)
}
