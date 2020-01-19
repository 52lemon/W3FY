package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"w3fy/controller/api/v1/Comment"
	"w3fy/controller/api/v1/Likes/TagLikes"
	"w3fy/controller/api/v1/Likes/TopicLikes"
	"w3fy/controller/api/v1/Relation"
	"w3fy/controller/api/v1/Topic"
	"w3fy/controller/api/v1/User"
	"w3fy/controller/api/v1/captcha"
	"w3fy/middleware"
	"w3fy/pkg/setting"
)

func InitRoute() *gin.Engine {
	gin.DisableConsoleColor()
	gin.SetMode(setting.RUNMODE)
	r := gin.New()
	//全局中间件
	r.Use(gin.Logger())
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
		//用户api
		user := apiv1.Group("/user")
		{
			user.POST("/register/c1", User.RegisterByUsername)
			user.POST("/login/c1", User.LoginByPassword)

			apiv1.GET("/userinfo", middleware.JWT(), User.GetUserInfo)
			apiv1.PUT("/userinfo", middleware.JWT(), User.UpdateUserInfo)
		}
		//验证码api
		capt := apiv1.Group("/captcha")
		{
			capt.GET("/create", captcha.GetCaptcha)
			capt.GET("/show/:source", captcha.ShowCaptcha)
			capt.GET("/reload/:source", captcha.ReloadCaptcha)
		}
		//帖子接口
		top := apiv1.Group("/topic")
		{
			top.POST("/", middleware.JWT(), Topic.CreateTopic)
			top.PUT("/", middleware.JWT(), Topic.UpdateTopic)
			top.DELETE("/", middleware.JWT(), Topic.DeleteTopic)
			top.GET("/", middleware.JWT(), Topic.GetUserTopics)

			top.GET("/normal", Topic.GetTopics)
			top.GET("/normal/:id", Topic.GetSingleTopic)

			top.GET("/search/type1", Topic.TagTopics)
			top.GET("/search/type2", Topic.TitleTopics)
		}
		//评论接口
		comm := apiv1.Group("/comment")
		{
			comm.GET("/:top_id", middleware.JWT(), Comment.GetComments)
			comm.POST("/", middleware.JWT(), Comment.CreateComment)
			comm.DELETE("/", middleware.JWT(), Comment.DeleteComment)
		}
		//节点收藏接口
		tag := apiv1.Group("/taglikes")
		{
			tag.POST("/", middleware.JWT(), TagLikes.CreateTagLike)
			tag.DELETE("/", middleware.JWT(), TagLikes.DeleteTagLike)
			tag.GET("/", middleware.JWT(), TagLikes.GetTagLikes)
		}
		//帖子收藏接口
		topl := apiv1.Group("/topiclikes")
		{
			topl.POST("/", middleware.JWT(), TopicLikes.GetTopicLikes)
			topl.DELETE("/", middleware.JWT(), TopicLikes.DeleteTopicLike)
			topl.GET("/", middleware.JWT(), TopicLikes.GetTopicLikes)
		}
		//关系接口
		rel := apiv1.Group("/relation")
		{
			rel.POST("/", middleware.JWT(), Relation.CreateRelation)
			rel.GET("/follow", middleware.JWT(), Relation.GetFollow)
			rel.GET("/follower", middleware.JWT(), Relation.GetFollower)
			rel.POST("/delete", middleware.JWT(), Relation.DeleteRelation)
		}

	}
	return r
}
