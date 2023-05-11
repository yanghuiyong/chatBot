package logic

import (
	redisPool "chatBot/until/redis"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"math/rand"
	"strconv"
	"time"
)

const (
	RedisChatBotApiKey = "redis:chat:bot:key"
	OnceApiKeyTime     = "chatApi:key:once:time:%s"
)

// GetChatApiKey ...
func GetChatApiKey() (string, error) {
	con := redisPool.RedisSrv.Get()
	defer con.Close()

	list, err := redis.Strings(con.Do("LRANGE", RedisChatBotApiKey, 0, 100))
	if err != nil {
		return "", err
	}
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(len(list))
	apiKey := list[num]
	// 先查
	kk := fmt.Sprintf(OnceApiKeyTime, apiKey)
	times, _ := redis.String(con.Do("get", kk))
	timeInt, _ := strconv.Atoi(times)
	if timeInt > 2 {
		rand.Seed(time.Now().UnixNano())
		num = rand.Intn(len(list))
		apiKey = list[num]
		kk = fmt.Sprintf(OnceApiKeyTime, apiKey)
	}

	con.Do("incr", kk, 1)
	if timeInt == 0 {
		con.Do("EXPIRE", kk, 70)
	}
	
	return apiKey, nil
}
