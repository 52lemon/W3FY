package Auth

import (
	"w3fy/models"
)

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (Auth) TableName() string {
	return "users"
}

func CheckAndReturnIdc1(username, password string) int {
	var auth Auth
	models.DB.Debug().Select("id").Where(Auth{Username: username, Password: password}).First(&auth)
	return int(auth.ID)
}
