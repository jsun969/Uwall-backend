package main

import (
	"Uwall/api"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/love", api.PostLove)
	r.Run()
}
