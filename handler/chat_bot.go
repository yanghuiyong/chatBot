package handler

import (
	"chatBot/until/zaplog"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

var RandomSliceApiKey = []string{
	"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",
	"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",
	"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",
	"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",
}

func GetChatReplay(c *gin.Context) {
	jData := map[string]interface{}{
		"ret":   0,
		"value": c.Query("aa"),
	}
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(4)
	apiKeys := RandomSliceApiKey[num]
	client := openai.NewClient(apiKeys)
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
	//fmt.Println(resp.Choices[0].Message.Content)
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
