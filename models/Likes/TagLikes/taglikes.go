package TagLikes

import (
	"w3fy/models"
	"w3fy/pkg/logging"
)

type TagLikes struct {
	Uid   int    `json:"uid"`
	Tname string `json:"tname"`
}

func (TagLikes) TableName() string {
	return "taglikes"
}

//获取用户收藏的节点
//select * from `taglikes` where(`uid`=xxx)
func GetTagLikes(uid int) (tags []TagLikes, err error) {
	if err = models.DB.Debug().Where("uid=?", uid).Find(&tags).Error; err != nil {
		logging.DebugLog(err)
	}
	return
}

//添加用户收藏的节点
//insert into `taglikes` values(xx)
func CreateTagLikes(likes TagLikes) bool {
	if err := models.DB.Debug().Create(&likes).Error; err != nil {
		logging.DebugLog(err)
		return false
	}
	return true
}

//删除用户收藏的节点
//delete from `taglikes` where(`uid`=xx,`tname`=xx)
func DeleteTagLikes(uid int, tname string) bool {
	if err := models.DB.Debug().Unscoped().Where(map[string]interface{}{"uid": uid, "tname": tname}).Delete(&TagLikes{}).Error; err != nil {
		logging.DebugLog(err)
		return false
	}
	return true
}
