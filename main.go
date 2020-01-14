package main

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"w3fy/middleware"
	"w3fy/models/User"
)

func main() {
	// 禁用控制台颜色, 将日志写入文件时不需要控制台颜色
	gin.DisableConsoleColor()
	router := gin.Default()
	router.Use(middleware.LoggerToFile())
	router.Use(gin.Recovery())
	router.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		user := User.GetInfo(com.StrTo(id).MustInt())
		c.JSON(200, &user)

	})
	router.Run(":8080")
}
