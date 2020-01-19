package Topic

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"strings"
	"w3fy/models/Topic"
	"w3fy/pkg/e"
	"w3fy/pkg/logging"
	"w3fy/pkg/setting"
	"w3fy/pkg/util"
)

//创建帖子
func CreateTopic(c *gin.Context) {
	code := e.INTERNAL_SERVER_ERROR
	data := make(map[string]interface{})
	var msg string

	//获取post数据
	tag := c.PostFormArray("tag")
	uid := c.MustGet("AuthData").(*util.Claims).User.ID
	title := c.PostForm("title")
	content := c.PostForm("content")

	//验证表单
	valid := validation.Validation{}
	valid.Required(tag, "tag").Message("tag不能为空")
	valid.MaxSize(tag, 3, "tag").Message("tag信息过多")
	valid.Required(title, "title").Message("title不能为空")
	valid.MaxSize(title, 120, "title").Message("title信息过长")
	valid.MaxSize(content, 20000, "content").Message("content信息过长")

	//若表单信息有误
	if valid.HasErrors() {
		code = e.BAD_REQUEST
		errordata := make(map[int]interface{})
		for index, err := range valid.Errors {
			logging.DebugLog(err.Key, err.Message)
			errordata[index] = map[string]interface{}{err.Key: err.Message}
		}
		data["error"] = errordata
	}
	//若表单信息无误
	if _, ok := data["error"]; !ok {
		//插入数据
		topic := Topic.Topic{Uid: uid, Tag: strings.Join(tag, ","), Title: title, Content: content}
		if Topic.CreateTopic(&topic) {
			code = e.OK
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
	msg = "表单数据异常"
	c.JSON(code, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
	return
}

//修改帖子内容
func UpdateTopic(c *gin.Context) {
	code := e.INTERNAL_SERVER_ERROR
	data := make(map[string]interface{})
	var msg string

	//获取post数据
	id := c.PostForm("id") //文章id
	uid := c.MustGet("AuthData").(*util.Claims).User.ID
	content := c.PostForm("content")

	//验证表单
	valid := validation.Validation{}
	valid.Required(id, "id").Message("id不能为空")
	valid.MaxSize(content, 20000, "content").Message("content信息过长")

	//若表单信息有误
	if valid.HasErrors() {
		code = e.BAD_REQUEST
		errordata := make(map[int]interface{})
		for index, err := range valid.Errors {
			logging.DebugLog(err.Key, err.Message)
			errordata[index] = map[string]interface{}{err.Key: err.Message}
		}
		data["error"] = errordata
	}

	//若表单信息无误
	if _, ok := data["error"]; !ok {
		//更新数据
		newid := com.StrTo(id).MustInt()
		if Topic.GetUidById(newid).Uid == uid {
			updata := map[string]interface{}{"content": content}
			if Topic.UpdateTopic(newid, updata) {
				code = e.OK
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
	}
	msg = "表单数据异常"
	c.JSON(code, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
	return
}

//删除帖子
func DeleteTopic(c *gin.Context) {
	code := e.INTERNAL_SERVER_ERROR
	data := make(map[string]interface{})
	var msg string

	//获取post数据
	id := c.Query("id")
	uid := c.MustGet("AuthData").(*util.Claims).User.ID

	//验证表单数据
	valid := validation.Validation{}
	valid.Required(id, "id").Message("id不能为空")

	//若表单信息有误
	if valid.HasErrors() {
		code = e.BAD_REQUEST
		errordata := make(map[int]interface{})
		for index, err := range valid.Errors {
			logging.DebugLog(err.Key, err.Message)
			errordata[index] = map[string]interface{}{err.Key: err.Message}
		}
		data["error"] = errordata
	}
	if _, ok := data["error"]; !ok {
		newid := com.StrTo(id).MustInt()
		if Topic.GetUidById(newid).Uid == uid {
			if Topic.DeleteTopic(newid) {
				code = e.OK
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
	}
	msg = "表单数据异常"
	c.JSON(code, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
	return
}

//用户查看帖子
func GetUserTopics(c *gin.Context) {
	code := e.INTERNAL_SERVER_ERROR
	data := make(map[string]interface{})
	var msg string

	userid := c.MustGet("AuthData").(*util.Claims).User.ID

	topics, err := Topic.GetUserTopic(userid)
	if err != nil {
		msg = "服务器异常"
		c.JSON(code, gin.H{
			"code": code,
			"data": data,
			"msg":  msg,
		})
		return
	}
	code = e.OK
	msg = "请求成功"
	data["user"] = topics
	c.JSON(code, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

//获取指定帖子数据
func GetSingleTopic(c *gin.Context) {
	code := e.INTERNAL_SERVER_ERROR
	data := make(map[string]interface{})
	var msg string

	id := c.Param("id")
	newid := com.StrTo(id).MustInt()
	topic, err := Topic.GetTopicById(newid)
	if err != nil {
		msg = "服务器异常"
		c.JSON(code, gin.H{
			"code": code,
			"data": data,
			"msg":  msg,
		})
		return
	}
	code = e.OK
	data["topic"] = topic
	msg = "请求成功"
	c.JSON(code, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

//分页获取帖子
func GetTopics(c *gin.Context) {
	code := e.INTERNAL_SERVER_ERROR
	data := make(map[string]interface{})
	var msg string

	topics, err := Topic.GetTopic(util.GetPage(c), setting.PAGE_SIZE)
	if err != nil {
		msg = "服务器异常"
		c.JSON(code, gin.H{
			"code": code,
			"data": data,
			"msg":  msg,
		})
		return
	}
	code = e.OK
	data["topics"] = topics
	msg = "请求成功"
	c.JSON(code, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

//根据标签查询帖子
func TagTopics(c *gin.Context) {
	code := e.INTERNAL_SERVER_ERROR
	data := make(map[string]interface{})
	var msg string

	tag := c.Query("tag")
	topics, err := Topic.SearchInTag(util.GetPage(c), setting.PAGE_SIZE, tag)
	if err != nil {
		msg = "服务器异常"
		c.JSON(code, gin.H{
			"code": code,
			"data": data,
			"msg":  msg,
		})
		return
	}
	code = e.OK
	data["topics"] = topics
	msg = "请求成功"
	c.JSON(code, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

//根据标题查询帖子
func TitleTopics(c *gin.Context) {
	code := e.INTERNAL_SERVER_ERROR
	data := make(map[string]interface{})
	var msg string

	title := c.Query("title")
	topics, err := Topic.SearchInTitle(util.GetPage(c), setting.PAGE_SIZE, title)
	if err != nil {
		msg = "服务器异常"
		c.JSON(code, gin.H{
			"code": code,
			"data": data,
			"msg":  msg,
		})
		return
	}
	code = e.OK
	data["topics"] = topics
	msg = "请求成功"
	c.JSON(code, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}
