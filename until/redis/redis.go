package redisPool

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

var (
	RedisSrv *redis.Pool
)

func poolGet() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   500,
		MaxActive: 5000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "127.0.0.1:6379",
				redis.DialPassword(""),
				redis.DialConnectTimeout(10*time.Second),
				redis.DialReadTimeout(10*time.Second),
				redis.DialWriteTimeout(10*time.Second))
			if err != nil {
				fmt.Println("redis Pool error: ", err.Error())
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
		IdleTimeout: 300 * time.Second,
		Wait:        true,
	}
}

// Init 初始化redis
func Init() {
	RedisSrv = poolGet()
}
