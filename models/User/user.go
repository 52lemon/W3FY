package User

import (
	"w3fy/models"
	"w3fy/pkg/logging"
)

type User struct {
	models.Model

	Username     string `json:"username"`
	Nickname     string `json:"nickname"`
	Password     string `json:"-"`
	Sex          int    `json:"sex"`
	Email        int    `json:"email"`
	Website      string `json:"website"`
	Education    string `json:"education"`
	Collage      string `json:"collage"`
	Introduction string `json:"introduction"`
	Github       string `json:"github"`
	Avatar       string `json:"avatar"`
	Coin         int    `json:"coin"`
}

func (User) TableName() string {
	return "users"
}

//获取用户信息
func GetInfo(id int) (user User) {
	err := models.DB.Debug().Where("id = ?", id).First(&user).Error
	if err != nil {
		logging.DebugLog(err)
	}
	return
}

//登录
func Login(user User) bool {
	var find User
	if models.DB.Where(user).Select("id").First(&find); find.ID > 0 {
		return true
	}
	return false
}
