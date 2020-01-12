package main

import (
	"github.com/gin-gonic/gin"
	"w3fy/middleware"
)

func main() {
	// 禁用控制台颜色, 将日志写入文件时不需要控制台颜色
	gin.DisableConsoleColor()
	router := gin.Default()
	router.Use(middleware.LoggerToFile())
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	router.Run(":8080")
}
