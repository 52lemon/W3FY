package Relation

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"w3fy/models/Relation"
	"w3fy/models/User"
	"w3fy/pkg/e"
	"w3fy/pkg/logging"
	"w3fy/pkg/setting"
	"w3fy/pkg/util"
)

//创建关系
func CreateRelation(c *gin.Context) {
	code := e.INTERNAL_SERVER_ERROR
	data := make(map[string]interface{})
	var msg string

	//获取post数据
	fromId := c.PostForm("from_id")
	toId := c.PostForm("to_id")
	userId := c.MustGet("AuthData").(*util.Claims).User.ID

	//验证表单
	valid := validation.Validation{}
	valid.Required(fromId, "from_id").Message("from_id不能为空")
	valid.Required(toId, "to_id").Message("to_id不能为空")
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
		//验证用户id的存在
		_, err := User.GetInfo(com.StrTo(toId).MustInt())
		if err == nil && userId == com.StrTo(fromId).MustInt() && fromId != toId { //不能自己关注自己
			relation := Relation.Relation{FromId: userId, ToId: com.StrTo(toId).MustInt()}
			if Relation.CreateRelation(relation) {
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
	}
	msg = "表单数据异常"
	c.JSON(code, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

//查看关注了谁
func GetFollow(c *gin.Context) {
	code := e.INTERNAL_SERVER_ERROR
	data := make(map[string]interface{})
	var msg string

	userId := c.MustGet("AuthData").(*util.Claims).User.ID
	follow, err := Relation.GetFocus(util.GetPage(c), setting.PAGE_SIZE, userId)
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
	data["follow"] = follow
	msg = "请求成功"
	c.JSON(code, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

//查看我的粉丝
func GetFollower(c *gin.Context) {
	code := e.INTERNAL_SERVER_ERROR
	data := make(map[string]interface{})
	var msg string

	userId := c.MustGet("AuthData").(*util.Claims).User.ID
	followers, err := Relation.GerFans(util.GetPage(c), setting.PAGE_SIZE, userId)
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
	data["followers"] = followers
	msg = "请求成功"
	c.JSON(code, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

//取消关注
func DeleteRelation(c *gin.Context) {
	code := e.INTERNAL_SERVER_ERROR
	data := make(map[string]interface{})
	var msg string

	//获取post数据
	fromId := c.PostForm("from_id")
	toId := c.PostForm("to_id")
	userId := c.MustGet("AuthData").(*util.Claims).User.ID

	//验证表单
	valid := validation.Validation{}
	valid.Required(fromId, "from_id").Message("from_id不能为空")
	valid.Required(toId, "to_id").Message("to_id不能为空")

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
		//验证用户id的存在
		_, err := User.GetInfo(com.StrTo(toId).MustInt())
		if err == nil && userId == com.StrTo(fromId).MustInt() {
			relation := Relation.Relation{FromId: userId, ToId: com.StrTo(toId).MustInt()}
			if Relation.DeleteRelation(&relation) {
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
}
