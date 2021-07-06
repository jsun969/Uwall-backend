package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostLove(c *gin.Context) {
	c.Status(http.StatusOK)
}
