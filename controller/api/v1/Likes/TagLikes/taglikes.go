package TagLikes

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"w3fy/models/Likes/TagLikes"
	"w3fy/pkg/e"
	"w3fy/pkg/logging"
	"w3fy/pkg/util"
)

//添加节点收藏
func CreateTagLike(c *gin.Context) {
	code := e.INTERNAL_SERVER_ERROR
	data := make(map[string]interface{})
	var msg string
	//获取post数据
	tname := c.PostForm("tname")
	userId := c.MustGet("AuthData").(*util.Claims).User.ID
	//验证表单
	valid := validation.Validation{}
	valid.Required(tname, "tname").Message("tname不能为空")
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
		taglike := TagLikes.TagLikes{Uid: userId, Tname: tname}
		if TagLikes.CreateTagLikes(taglike) {
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

//删除节点收藏
func DeleteTagLike(c *gin.Context) {
	code := e.INTERNAL_SERVER_ERROR
	data := make(map[string]interface{})
	var msg string
	//获取post数据
	tname := c.PostForm("tname")
	userId := c.MustGet("AuthData").(*util.Claims).User.ID
	//验证表单
	valid := validation.Validation{}
	valid.Required(tname, "tname").Message("tname不能为空")
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
		taglike := TagLikes.TagLikes{Uid: userId, Tname: tname}
		if TagLikes.DeleteTagLikes(&taglike) {
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

//查看收藏节点
func GetTagLikes(c *gin.Context) {
	code := e.INTERNAL_SERVER_ERROR
	data := make(map[string]interface{})
	var msg string

	userId := c.MustGet("AuthData").(*util.Claims).User.ID
	taglikes, err := TagLikes.GetTagLikes(userId)
	if err != nil {
		msg = "服务器异常"
		c.JSON(code, gin.H{
			"code": code,
			"data": data,
			"msg":  msg,
		})
	}
	code = e.OK
	data["taglikes"] = taglikes
	msg = "请求成功"
	c.JSON(code, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}
