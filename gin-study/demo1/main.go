package main

import (
	"github.com/gin-gonic/gin"
	"demo1/bootstrap"
    "demo1/global"
)

func main() {
	bootstrap.InitializeConfig()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	
	  // 启动服务器
    r.Run(":" + global.App.Config.App.Port)
}
