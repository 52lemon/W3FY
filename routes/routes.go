package routes

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"w3fy/controller/api/v1/User"
	_ "w3fy/docs"
	"w3fy/middleware"
)

func InitRoute() *gin.Engine {
	gin.DisableConsoleColor()
	r := gin.New()
	r.Use(middleware.JWT())
	r.Use(middleware.LoggerToFile())
	r.Use(gin.Recovery())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"msg": "pong",
			})
		})
		apiv1.POST("/user/register/c1", User.RegisterByUsername)
	}
	return r
}
