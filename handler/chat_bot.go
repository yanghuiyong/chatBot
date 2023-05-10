package handler

import (
	"bytes"
	"chatBot/until/zaplog"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"net/http"
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
	/*rand.Seed(time.Now().UnixNano())
	num := rand.Intn(4)
	apiKeys := RandomSliceApiKey[num]*/
	/*client := openai.NewClient("sk-pKHZD1fLYqXDjjsdsdsdUvIODTT3ssjdfadsJC2gTuqqhTum")
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
	}*/

	msg, err := Completions(c.Query("question"))

	//fmt.Println(resp.Choices[0].Message.Content)
	zaplog.Trace("GetChatReplayRequest").Info("GetGpt", zap.Any("sReq", jData), zap.Any("err", err), zap.Error(err))
	if err != nil {
		jData["ret"] = 10052
		jData["errMsg"] = err.Error()
		c.JSONP(200, jData)
		return
	}
	jData["value"] = msg
	c.JSONP(200, jData)
	return
}

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatGPTRequestBody 请求体
type ChatGPTRequestBody struct {
	Model    string        `json:"model"`
	Messages []ChatMessage `json:"messages"`
}

type ResponseChoice struct {
	Index        int         `json:"index"`
	Message      ChatMessage `json:"message"`
	FinishReason string      `json:"finish_reason"`
}

type ResponseUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// ChatGPTResponseBody 响应体
type ChatGPTResponseBody struct {
	ID      string           `json:"id"`
	Object  string           `json:"object"`
	Created int              `json:"created"`
	Choices []ResponseChoice `json:"choices"`
	Usage   ResponseUsage    `json:"usage"`
}

type Context struct {
	Request  string
	Response string
	Time     int64
}

type ContextMgr struct {
	contextList []*Context
}

func (m *ContextMgr) Init() {
	m.contextList = make([]*Context, 10)
}

func (m *ContextMgr) checkExpire() {
	timeNow := time.Now().Unix()
	if len(m.contextList) > 0 {
		startPos := len(m.contextList) - 1
		for i := 0; i < len(m.contextList); i++ {
			if timeNow-m.contextList[i].Time < 1*60 {
				startPos = i
				break
			}
		}

		m.contextList = m.contextList[startPos:]
	}
}

func (m *ContextMgr) AppendMsg(request string, response string) {
	m.checkExpire()
	context := &Context{Request: request, Response: response, Time: time.Now().Unix()}
	m.contextList = append(m.contextList, context)
}

func (m *ContextMgr) GetData() []*Context {
	m.checkExpire()
	return m.contextList
}

var contextMgr ContextMgr

type ChatGPTErrorBody struct {
	Error map[string]interface{} `json:"error"`
}

func Completions(msg string) (string, error) {
	apiKey := "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"
	var messages []ChatMessage
	messages = append(messages, ChatMessage{
		Role:    "system",
		Content: "You are a helpful assistant.",
	})

	list := contextMgr.GetData()
	for i := 0; i < len(list); i++ {
		messages = append(messages, ChatMessage{
			Role:    "user",
			Content: list[i].Request,
		})

		messages = append(messages, ChatMessage{
			Role:    "assistant",
			Content: list[i].Response,
		})
	}

	messages = append(messages, ChatMessage{
		Role:    "user",
		Content: msg,
	})

	requestBody := ChatGPTRequestBody{
		Model:    "gpt-3.5-turbo",
		Messages: messages,
	}
	requestData, err := json.Marshal(requestBody)

	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(requestData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(response.Body)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	gptResponseBody := &ChatGPTResponseBody{}
	err = json.Unmarshal(body, gptResponseBody)
	if err != nil {
		return "", err
	}

	var reply string
	if len(gptResponseBody.Choices) > 0 {
		for _, v := range gptResponseBody.Choices {
			reply += "\n"
			reply += v.Message.Content
		}

		contextMgr.AppendMsg(msg, reply)
	}

	if len(reply) == 0 {
		gptErrorBody := &ChatGPTErrorBody{}
		err = json.Unmarshal(body, gptErrorBody)
		if err != nil {
			return "", err
		}

		reply += "Error: "
		reply += gptErrorBody.Error["message"].(string)
	}

	return reply, nil
}
