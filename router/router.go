package router

import (
	"chatBot/handler"
	"github.com/gin-gonic/gin"
)

func InitRouter(e *gin.Engine) {
	api := e.Group("/api")
	{
		api.GET("getChatReplay", handler.GetChatReplay)
	}
}
