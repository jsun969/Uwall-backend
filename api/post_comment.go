package api

import (
	"Uwall/database"
	"Uwall/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostComment(c *gin.Context) {
	var req model.ReqComment
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	comment := model.Comment{
		MessageID: req.ID,
		Name:      req.Name,
		Comment:   req.Comment,
		Anonymous: req.Anonymous,
	}
	if result := database.Db.Create(&comment); result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": result.Error})
		return
	}
	c.Status(http.StatusCreated)
}
