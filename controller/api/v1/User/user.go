package User

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"strconv"
	"w3fy/models/Auth"
	"w3fy/models/User"
	"w3fy/pkg/e"
	"w3fy/pkg/logging"
	"w3fy/pkg/util"
)

func RegisterByUsername(c *gin.Context) {
	//定义数据格式
	data := make(map[string]interface{})
	code := e.INTERNAL_SERVER_ERROR
	var msg string

	//获取post数据
	username := c.PostForm("username")
	password := c.PostForm("password")
	id := c.Param("id")
	digits := c.Param("digits")

	//请求数据处理
	valid := validation.Validation{}
	valid.Required(username, "username").Message("用户名不能为空")
	valid.MinSize(username, 6, "username").Message("用户名过短")
	valid.MaxSize(username, 20, "username").Message("用户名过长")
	valid.Required(password, "password").Message("密码不能为空")
	valid.MinSize(password, 6, "username").Message("密码过短")
	valid.MaxSize(password, 20, "username").Message("密码过长")
	valid.Required(id, "id").Message("验证码id不能为空")
	valid.Required(digits, "digits").Message("验证码不能为空")
	//若表单验证不成功,处理异常请求
	if valid.HasErrors() {
		code = e.BAD_REQUEST
		errorData := make(map[string]interface{})
		for index, err := range valid.Errors {
			logging.DebugLog(err.Key, err.Message)
			errorData[strconv.Itoa(index)] = map[string]interface{}{err.Key: err.Message}
		}
		data["error"] = errorData
	}
	//若表单验证成功
	if _, ok := data["error"]; !ok {
		//验证码校验
		if captcha.VerifyString(id, digits) {
			//密码md5加密
			md5Ctx := md5.New()
			md5Ctx.Write([]byte(password))
			password = fmt.Sprintf("%x", md5Ctx.Sum(nil))
			//用户账号入库
			user := User.User{
				Username: username,
				Password: password,
			}
			if User.CreateByPassword(&user) {
				id := Auth.CheckAndReturnIdc1(username, password)
				if id > 0 {
					token, err := util.GenerateToken(id)
					if err == nil {
						code = e.CREATED
						data["token"] = token
					}
				} else {
					code = e.UNAUTHORIZED
					msg = "用户名或密码错误"
				}
			}
			if msg == "" {
				msg = e.GetMsg(code)
			}
			c.JSON(code, gin.H{
				"code": code,
				"data": data,
				"msg":  msg,
			})
		} else {
			code = e.BAD_REQUEST
			msg = "验证码错误"
			c.JSON(code, gin.H{
				"code": code,
				"data": data,
				"msg":  msg,
			})
		}
	}
}
