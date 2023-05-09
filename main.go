package main

import (
	"chatBot/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	engine := gin.New()
	engine.Use(cors.New(cors.Config{
		//AllowOrigins:     []string{sysCfg["host"]},
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "x-requested-with", "content-type", "Clsq-Verify"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.InitRouter(engine)
	//用于健康检查
	engine.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"code": 200, "message": "ok"})
	})

	s := &http.Server{
		Addr:           ":8088",
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
