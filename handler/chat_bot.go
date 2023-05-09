package handler

import (
	"chatBot/until/zaplog"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetChatReplay(c *gin.Context) {
	jData := map[string]interface{}{
		"ret":   0,
		"value": c.Query("aa"),
	}
	zaplog.Trace("GetChatReplayRequest").Info("", zap.Any("sReq", jData))
	c.JSONP(200, jData)
	return
}
