package User

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"regexp"
	"strconv"
	"w3fy/models/Auth"
	"w3fy/models/User"
	"w3fy/pkg/e"
	"w3fy/pkg/logging"
	"w3fy/pkg/util"
)

//用户注册
func RegisterByUsername(c *gin.Context) {
	//定义数据格式
	data := make(map[string]interface{})
	code := e.INTERNAL_SERVER_ERROR
	var msg string

	//获取post数据
	username := c.PostForm("username")
	password := c.PostForm("password")
	id := c.PostForm("VerificationId")
	digits := c.PostForm("VerificationCode")

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
						msg = "注册成功"
						c.JSON(code, gin.H{
							"code": code,
							"data": data,
							"msg":  msg,
						})
						return
					}
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
		} else {
			code = e.BAD_REQUEST
			msg = "验证码错误"
			c.JSON(code, gin.H{
				"code": code,
				"data": data,
				"msg":  msg,
			})
			return
		}
	}
	//若表单验证不成功
	msg = "表单输入异常"
	c.JSON(code, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
	return
}

//用户登录
func LoginByPassword(c *gin.Context) {
	//定义数据格式
	code := e.INTERNAL_SERVER_ERROR
	data := make(map[string]interface{})
	var msg string

	//获取post数据
	username := c.PostForm("username")
	password := c.PostForm("password")
	id := c.PostForm("VerificationId")
	digits := c.PostForm("VerificationCode")

	//验证表单
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
		errordata := make(map[int]interface{})
		for index, err := range valid.Errors {
			logging.DebugLog(err.Key, err.Message)
			errordata[index] = map[string]interface{}{err.Key: err.Message}
		}
		data["error"] = errordata
	}
	//若表单验证成功
	if _, ok := data["error"]; !ok {
		//验证验证码
		if captcha.VerifyString(id, digits) {
			//验证账号
			id := Auth.CheckAndReturnIdc1(username, password)
			if id > 0 {
				token, err := util.GenerateToken(id)
				if err == nil {
					code = e.CREATED
					data["token"] = token
					msg = e.GetMsg(code)
					c.JSON(code, gin.H{
						"code": code,
						"data": data,
						"msg":  msg,
					})
					return
				}
			} else {
				msg = "服务器异常"
				c.JSON(code, gin.H{
					"code": code,
					"data": data,
					"msg":  msg,
				})
				return
			}
		} else {
			code = e.UNAUTHORIZED
			msg = "验证码错误"
			c.JSON(code, gin.H{
				"code": code,
				"data": data,
				"msg":  msg,
			})
			return
		}
	}
	//若表单验证不成功
	msg = "表单输入异常"
	c.JSON(code, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
	return
}

//浏览个人信息
func GetUserInfo(c *gin.Context) {
	code := e.OK
	data := make(map[string]interface{})
	var msg string
	//解析token
	userId := c.MustGet("AuthData").(*util.Claims).User.ID
	user, err := User.GetInfo(userId)
	//数据查询失败
	if err != nil {
		code = e.INTERNAL_SERVER_ERROR
		msg = "服务器异常"
		c.JSON(code, gin.H{
			"code": code,
			"data": data,
			"msg":  msg,
		})
		return
	}
	data["userInfo"] = user
	msg = "请求成功"
	c.JSON(code, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

//修改个人信息
func UpdateUserInfo(c *gin.Context) {
	code := e.OK
	data := make(map[string]interface{})
	var msg string
	reg := regexp.MustCompile(`(http|ftp|https):\/\/[\w\-_]+(\.[\w\-_]+)+([\w\-\.,@?^=%&:/~\+#]*[\w\-\@?^=%&/~\+#])?`)

	//解析token
	userId := c.MustGet("AuthData").(*util.Claims).User.ID
	nickname := c.PostForm("nickname")
	email := c.PostForm("email")
	website := c.PostForm("website")
	sex := c.PostForm("sex")
	collage := c.PostForm("collage")
	introduction := c.PostForm("introduction")
	github := c.PostForm("github")
	avatar := c.PostForm("avatar")
	education := c.PostForm("education")
	//验证表单
	valid := validation.Validation{}
	//nickname
	if nickname != "" {
		valid.MaxSize(nickname, 10, "nickname").Message("nickname信息过长")
	}
	//email
	if email != "" {
		valid.Email(email, "email").Message("email格式不正确")
	}
	//website
	if website != "" {
		valid.Match(website, reg, "website").Message("website格式不正确")
	}
	//sex
	if sex != "" {
		valid.Range(com.StrTo(sex).MustInt(), 0, 2, "sex").Message("性别异常")
	}
	//collage
	if collage != "" {
		valid.MaxSize(collage, 20, "collage").Message("collage信息过长")
	}
	//introduction
	if introduction != "" {
		valid.MaxSize(introduction, 80, "collage").Message("collage信息过长")
	}
	//github
	if github != "" {
		valid.Match(github, reg, "website").Message("github格式不正确")
	}
	//avatar
	if avatar != "" {
		valid.Match(avatar, reg, "avatar").Message("avatar格式不正确")
	}
	//education
	if education != "" {
		valid.MaxSize(education, 3, "education").Message("education信息过长")
	}
	//若表单有错误
	if valid.HasErrors() {
		code = e.BAD_REQUEST
		errordata := make(map[int]interface{})
		for index, err := range valid.Errors {
			logging.DebugLog(err.Key, err.Message)
			errordata[index] = map[string]interface{}{err.Key: err.Message}
		}
		data["error"] = errordata
	}
	//若表单验证成功
	if _, ok := data["error"]; !ok {
		user, err := User.GetInfo(userId)
		if err != nil {
			code = e.INTERNAL_SERVER_ERROR
			msg = "服务器异常"
			c.JSON(code, gin.H{
				"code": code,
				"data": data,
				"msg":  msg,
			})
			return
		}
		updatedata := map[string]interface{}{"nickname": nickname, "email": email, "website": website, "sex": com.StrTo(sex).MustInt(),
			"collage": collage, "introduction": introduction, "github": github, "avatar": avatar, "education": education}
		if User.UpdateInfo(&user, updatedata) {
			msg = "请求成功"
			c.JSON(code, gin.H{
				"code": code,
				"data": data,
				"msg":  msg,
			})
			return
		} else {
			code = e.INTERNAL_SERVER_ERROR
			msg = "服务器异常"
			c.JSON(code, gin.H{
				"code": code,
				"data": data,
				"msg":  msg,
			})
			return
		}
	}
	//若表单验证不成功
	msg = "表单输入异常"
	c.JSON(code, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
	return
}
