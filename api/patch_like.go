package api

import (
	"Uwall/database"
	"Uwall/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func PatchLike(c *gin.Context) {
	var req model.ReqLikeId
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if result := database.Db.Model(&model.Message{}).Where("id = ?", req.ID).Update("likes", gorm.Expr("likes + ?", 1)); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}
	c.Status(http.StatusOK)
}
