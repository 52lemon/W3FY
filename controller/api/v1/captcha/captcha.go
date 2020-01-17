package captcha

import (
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"time"
	"w3fy/pkg/e"
	"w3fy/pkg/logging"
)

//获取验证码
func GetCaptcha(c *gin.Context) {
	code := e.INTERNAL_SERVER_ERROR
	var msg string

	ID := captcha.NewLen(4)
	c.Set("id", ID) //在上下文中设置变量
	err := captcha.WriteImage(c.Writer, ID, captcha.StdWidth, captcha.StdHeight)
	if err != nil {
		logging.DebugLog(err)
		msg = "fail to get captcha"
		c.JSON(code, gin.H{
			"code": code,
			"msg":  msg,
			"data": "",
		})
		return
	}
	code = e.OK
	msg = "success"
	c.JSON(code, gin.H{
		"code": code,
		"msg":  msg,
		"data": map[string]interface{}{"id": ID, "time": time.Now().Unix()},
	})
}

//重载验证码
func ReloadCaptcha(c *gin.Context) {
	code := e.INTERNAL_SERVER_ERROR
	var msg string

	ID, isExist := c.Get("id")
	if isExist {
		if captcha.Reload(ID.(string)) {
			err := captcha.WriteImage(c.Writer, ID.(string), captcha.StdWidth, captcha.StdHeight)
			if err != nil {
				logging.DebugLog(err)
			}
			code = e.OK
			msg = "success"
			c.JSON(200, gin.H{
				"code": code,
				"msg":  msg,
				"data": map[string]interface{}{"id": ID.(string), "time": time.Now().Unix()},
			})
			return
		}
	}
	msg = "fail to update captcha"
	c.JSON(code, gin.H{
		"code": code,
		"msg":  msg,
		"data": "",
	})
}
