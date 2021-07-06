package api

import (
	"Uwall/database"
	"Uwall/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostLove(c *gin.Context) {
	var req model.Love
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	message := model.Message{
		FromName:  req.From.Name,
		FromSex:   req.From.Sex,
		ToName:    req.To.Name,
		ToSex:     req.To.Sex,
		Message:   req.Message,
		Anonymous: req.Anonymous,
	}
	if result := database.Db.Create(&message); result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": result.Error})
	}
	c.Status(http.StatusCreated)
}
