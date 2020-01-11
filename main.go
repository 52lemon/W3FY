package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"w3fy/pkg/logging"
)

func main() {
	// 禁用控制台颜色, 将日志写入文件时不需要控制台颜色
	gin.DisableConsoleColor()

	// 日志写入文件
	gin.DefaultWriter = io.MultiWriter(logging.SF)

	// 如果需要同时将日志写入文件和控制台, 请使用以下代码
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.Run(":8080")
}
