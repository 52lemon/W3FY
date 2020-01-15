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
	router.Use(middleware.LoggerToFile())
	router.Use(gin.Recovery())
	router.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		user := User.GetInfo(com.StrTo(id).MustInt())
		c.JSON(200, &user)

	})
	//model测试用例
	router.GET("/testcase01/:id2", func(c *gin.Context) {
		id := c.Param("id2")
		if Tags.CreateTag(&Tags.Tags{Name: id}) {
			c.JSON(200, map[string]interface{}{"msg": "success"})
		} else {
			c.JSON(500, map[string]interface{}{"msg": "failed"})
		}

	})
	router.GET("/testcase02", func(c *gin.Context) {
		tags := Tags.GetTags()
		c.JSON(200, &tags)
	})
	router.GET("/testcase03/:id3", func(c *gin.Context) {
		id := c.Param("id3")
		if Tags.DeleteTag(id) {
			c.JSON(200, map[string]interface{}{"msg": "success"})
		} else {
			c.JSON(500, map[string]interface{}{"msg": "failed"})
		}
	})
	router.Run(":8080")
}
