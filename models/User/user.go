package User

import (
	"w3fy/models"
	"w3fy/pkg/logging"
)

//定义用户模型
type User struct {
	models.Model

	Username     string `json:"username"`
	Nickname     string `json:"nickname"`
	Password     string `json:"-"`
	Sex          int    `json:"sex"`
	Email        string `json:"email"`
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

//用户名注册 insert into user (`username`,`password`) values(xxx,xxx)
func CreateByPassword(user *User) bool {
	if models.DB.NewRecord(user) {
		models.DB.Debug().Create(user)
		return !models.DB.NewRecord(user)
	}
	return false
}

//登录  select id from user   check user.ID
func Login(user User) bool {
	var find User
	if models.DB.Where(user).Select("id").First(&find); find.ID > 0 {
		return true
	}
	return false
}

//获取用户信息 select * from user where id = ?
func GetInfo(id int) (user User) {
	err := models.DB.Debug().Where("id = ?", id).First(&user).Error
	if err != nil {
		logging.DebugLog(err)
	}
	return
}

//更新用户信息 update user set `aa` = xx ,`bb`=xx
func UpdateInfo(user *User, data interface{}) bool {
	if err := models.DB.Debug().Model(user).Updates(data).Error; err == nil {
		return true
	}
	return false
}
