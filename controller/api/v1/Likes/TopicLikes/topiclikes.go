package TopicLikes

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"w3fy/models/Likes/TopicLikes"
	"w3fy/pkg/e"
	"w3fy/pkg/logging"
	"w3fy/pkg/util"
)

//添加帖子收藏
func CreateTopicTag(c *gin.Context) {
	code := e.INTERNAL_SERVER_ERROR
	data := make(map[string]interface{})
	var msg string

	//获取post数据
	tid := c.PostForm("tid")
	userId := c.MustGet("AuthData").(*util.Claims).User.ID
	//验证表单
	valid := validation.Validation{}
	valid.Required(tid, "tid").Message("tid不能为空")

	//若表单数据有误
	if valid.HasErrors() {
		code = e.BAD_REQUEST
		errordata := make(map[int]interface{})
		for index, err := range valid.Errors {
			logging.DebugLog(err.Key, err.Message)
			errordata[index] = map[string]interface{}{err.Key: err.Message}
		}
		data["error"] = errordata
	}
	//若表单数据无误
	if _, ok := data["error"]; !ok {
		//插入数据
		topiclike := TopicLikes.TopicLikes{Uid: userId, Tid: com.StrTo(tid).MustInt()}
		if TopicLikes.CreateTopicLikes(topiclike) {
			code = e.CREATED
			msg = "请求成功"
			c.JSON(code, gin.H{
				"code": code,
				"data": data,
				"msg":  msg,
			})
			return
		} else {
			msg = "服务器异常"
			c.JSON(code, gin.H{
				"code": code,
				"data": data,
				"msg":  msg,
			})
			return
		}
	}
	code = e.BAD_REQUEST
	msg = "表单数据异常"
	c.JSON(code, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

//删除帖子收藏
func DeleteTopicLike(c *gin.Context) {
	code := e.INTERNAL_SERVER_ERROR
	data := make(map[string]interface{})
	var msg string

	//获取post数据
	tid := c.Query("tid")
	userId := c.MustGet("AuthData").(*util.Claims).User.ID
	//验证表单
	valid := validation.Validation{}
	valid.Required(tid, "tid").Message("tid不能为空")

	//若表单数据有误
	if valid.HasErrors() {
		code = e.BAD_REQUEST
		errordata := make(map[int]interface{})
		for index, err := range valid.Errors {
			logging.DebugLog(err.Key, err.Message)
			errordata[index] = map[string]interface{}{err.Key: err.Message}
		}
		data["error"] = errordata
	}
	//若表单数据无误
	if _, ok := data["error"]; !ok {
		//删除数据
		if TopicLikes.DeleteTopicLikes(userId, com.StrTo(tid).MustInt()) {
			code = e.CREATED
			msg = "请求成功"
			c.JSON(code, gin.H{
				"code": code,
				"data": data,
				"msg":  msg,
			})
			return
		} else {
			msg = "服务器异常"
			c.JSON(code, gin.H{
				"code": code,
				"data": data,
				"msg":  msg,
			})
			return
		}
	}
	code = e.BAD_REQUEST
	msg = "表单数据异常"
	c.JSON(code, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

//查看帖子收藏
func GetTopicLikes(c *gin.Context) {
	code := e.INTERNAL_SERVER_ERROR
	data := make(map[string]interface{})
	var msg string

	userId := c.MustGet("AuthData").(*util.Claims).User.ID
	topiclikes, err := TopicLikes.GetTopicLikes(userId)
	if err != nil {
		msg = "服务器异常"
		c.JSON(code, gin.H{
			"code": code,
			"data": data,
			"msg":  msg,
		})
	}
	code = e.OK
	data["taglikes"] = topiclikes
	msg = "请求成功"
	c.JSON(code, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}
