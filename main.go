package main

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"w3fy/middleware"
	"w3fy/models/Tags"
	"w3fy/models/User"
)

func main() {
	// 禁用控制台颜色, 将日志写入文件时不需要控制台颜色
	gin.DisableConsoleColor()
	router := gin.Default()
	router.Use(middleware.JWT())
	router.Use(middleware.LoggerToFile())
	router.Use(gin.Recovery())
	router.GET("/usertest/:id", func(c *gin.Context) {
		id := c.Param("id")
		user := User.GetInfo(com.StrTo(id).MustInt())
		c.JSON(200, &user)

	})
	//model测试用例
	router.GET("/testcasev1/:name", func(c *gin.Context) {
		name := c.Param("name")
		if Tags.CreateTag(&Tags.Tags{Name: name}) {
			c.JSON(200, map[string]interface{}{"msg": "success"})
		} else {
			c.JSON(500, map[string]interface{}{"msg": "failed"})
		}

	})
	router.GET("/testcasev2", func(c *gin.Context) {
		tags := Tags.GetTags()
		c.JSON(200, &tags)
	})
	router.GET("/testcasev3/:key", func(c *gin.Context) {
		key := c.Param("key")
		if Tags.DeleteTag(key) {
			c.JSON(200, map[string]interface{}{"msg": "success"})
		} else {
			c.JSON(500, map[string]interface{}{"msg": "failed"})
		}
	})
	router.Run(":8080")
}
