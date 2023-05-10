package handler

import (
	"chatBot/until/zaplog"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	"go.uber.org/zap"
)

func GetChatReplay(c *gin.Context) {
	jData := map[string]interface{}{
		"ret":   0,
		"value": c.Query("aa"),
	}
	client := openai.NewClient("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")
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: c.Query("question"),
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	zaplog.Trace("GetChatReplayRequest").Info("GetGpt", zap.Any("sReq", jData), zap.Any("err", err), zap.Error(err))
	if err != nil {
		jData["ret"] = 10052
		jData["errMsg"] = err.Error()
		c.JSONP(200, jData)
		return
	}
	jData["value"] = resp.Choices[0].Message.Content
	c.JSONP(200, jData)
	return
}
