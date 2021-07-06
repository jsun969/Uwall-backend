package initialize

import (
	"Uwall/api"
	"github.com/gin-gonic/gin"
)

func Gin() {
	r := gin.Default()
	r.POST("/love", api.PostLove)
	r.Run()
}
