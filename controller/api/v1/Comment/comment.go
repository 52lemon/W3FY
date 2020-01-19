package Comment

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"w3fy/models/Comments"
	"w3fy/pkg/e"
	"w3fy/pkg/logging"
	"w3fy/pkg/setting"
	"w3fy/pkg/util"
)

//查看帖子的评论
func GetComments(c *gin.Context) {
	code := e.INTERNAL_SERVER_ERROR
	data := make(map[string]interface{})
	var msg string

	topId := c.Param("top_id")
	newtopId := com.StrTo(topId).MustInt()
	comments, err := Comments.GetCommentByTid(util.GetPage(c), setting.PAGE_SIZE, newtopId)
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
	data["comments"] = comments
	msg = "请求成功"
	c.JSON(code, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

//添加评论
func CreateComment(c *gin.Context) {
	code := e.INTERNAL_SERVER_ERROR
	data := make(map[string]interface{})
	var msg string

	//获取post数据
	topId := c.PostForm("top_id")
	fatherId := c.PostForm("father_id")
	fromId := c.MustGet("AuthData").(*util.Claims).User.ID
	toId := c.PostForm("to_id")
	comments := c.PostForm("comments")

	//验证表单数据
	valid := validation.Validation{}
	valid.Required(topId, "topId").Message("top_id不能为空")
	valid.Required(fatherId, "fatherId").Message("father_id不能为空")
	valid.Required(toId, "toId").Message("to_id不能为空")
	valid.Required(comments, "comments").Message("comments不能为空")
	valid.MaxSize(comments, 10000, "comments").Message("comments信息过长")

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
		comments := Comments.Comment{TopId: com.StrTo(topId).MustInt(), FromId: fromId, ToId: com.StrTo(toId).MustInt(), Comments: comments}
		if Comments.CreateComment(comments) {
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
	msg = "表单数据异常"
	c.JSON(code, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

//删除评论
func DeleteComment(c *gin.Context) {
	code := e.INTERNAL_SERVER_ERROR
	data := make(map[string]interface{})
	var msg string

	id := c.Query("id")
	newId := com.StrTo(id).MustInt()
	userId := c.MustGet("AuthData").(*util.Claims).User.ID

	if Comments.GetFromId(newId) == userId {
		if Comments.DeleteComment(newId) {
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
	msg = "接口请求异常"
	c.JSON(code, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}
