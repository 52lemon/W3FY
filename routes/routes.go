package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"w3fy/controller/api/v1/User"
	"w3fy/middleware"
	"w3fy/pkg/setting"
)

func InitRoute() *gin.Engine {
	gin.DisableConsoleColor()
	gin.SetMode(setting.RUNMODE)
	r := gin.New()
	//全局中间件
	r.Use(middleware.LoggerToFile())
	r.Use(cors.Default())
	r.Use(gin.Recovery())
	apiv1 := r.Group("/api/v1")
	{ //                             局部中间件
		apiv1.GET("ping", middleware.JWT(), func(c *gin.Context) {
			c.JSON(200, gin.H{
				"msg": "pong",
			})
		})
		apiv1.POST("/user/register/c1", User.RegisterByUsername)
	}
	return r
}
