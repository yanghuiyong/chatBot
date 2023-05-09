package handler

import (
	"chatBot/until/zaplog"
	"github.com/gin-gonic/gin"
	"github.com/solywsh/chatgpt"
	"go.uber.org/zap"
	"time"
)

func GetChatReplay(c *gin.Context) {
	jData := map[string]interface{}{
		"ret":   0,
		"value": c.Query("aa"),
	}

	chat := chatgpt.New(c.Query("appkey"), "zhangsan", 10*time.Second)
	defer chat.Close()
	answer, err := chat.ChatWithContext(c.Query("question"))
	if err != nil {
		jData["ret"] = 500
		jData["errMsg"] = err.Error()
		c.JSONP(200, jData)
		return
	}
	jData["value"] = answer
	zaplog.Trace("GetChatReplayRequest").Info("", zap.Any("sReq", jData))
	c.JSONP(200, jData)
	return
}
