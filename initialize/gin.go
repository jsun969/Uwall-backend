package initialize

import (
	"Uwall/api"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Gin() {
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/love", api.PostLove)
	r.POST("/complaint", api.PostComplaint)
	r.POST("/help", api.PostHelp)
	r.POST("/notice", api.PostNotice)
	r.POST("/expand", api.PostExpand)
	r.GET("/messages", api.GetMessages)
	r.GET("/messages/:type", api.GetMessagesWithType)
	r.POST("/comment", api.PostComment)
	r.PATCH("/like", api.PatchLike)
	r.Run(":1314")
}
