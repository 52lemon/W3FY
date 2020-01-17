package captcha

import (
	"fmt"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"time"
	"w3fy/pkg/e"
	"w3fy/pkg/logging"
	"w3fy/pkg/util"
)

// 用户---->GetCaptcha----(答错)------------->GetCaptcha
//            |(看不清楚验证码,更换验证码)
//          ReloadCaptcha---(答错)------------->GetCaptcha
//缺陷:恶意用户频繁请求GetCaptcha接口会造成服务器短期内存疯涨,需要控制用户请求频率
//获取验证码
func GetCaptcha(c *gin.Context) {
	code := e.INTERNAL_SERVER_ERROR
	var msg string
	//验证码长度为4个数字
	ID := captcha.NewLen(4)

	code = e.OK
	msg = "success"
	png := fmt.Sprintf("%s.png", ID)
	c.JSON(code, gin.H{
		"code": code,
		"msg":  msg,
		"data": map[string]interface{}{"png": png, "time": time.Now().Unix()},
	})
}

//展示验证码
func ShowCaptcha(c *gin.Context) {
	id := util.Ext(c.Param("source"))
	if err := captcha.WriteImage(c.Writer, id, captcha.StdWidth, captcha.StdHeight); err != nil {
		logging.DebugLog(err)
	}
	return
}

////重载验证码
func ReloadCaptcha(c *gin.Context) {
	id := util.Ext(c.Param("source"))
	if captcha.Reload(id) {
		err := captcha.WriteImage(c.Writer, id, captcha.StdWidth, captcha.StdHeight)
		if err != nil {
			logging.DebugLog(err)
		}
	}
}
