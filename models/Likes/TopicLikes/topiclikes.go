package TopicLikes

import (
	"w3fy/models"
	"w3fy/pkg/logging"
)

type TopicLikes struct {
	models.Model

	Uid int `json:"uid"`
	Tid int `json:"tid"`
}

func (TopicLikes) TableName() string {
	return "topiclikes"
}

//获取用户收藏的帖子
//select * from `topiclikes` where (`uid`=xx)
func GetTopicLikes(uid int) (topic []TopicLikes) {
	if err := models.DB.Debug().Where("uid=?", uid).Find(&topic).Error; err != nil {
		logging.DebugLog(err)
	}
	return
}

//删除用户收藏的帖子
//delete * from `topiclikes` where(`uid`=xx and `tid`=xx)
func DeleteTopicLikes(likes TopicLikes) bool {
	if err := models.DB.Debug().Unscoped().Delete(&likes).Error; err != nil {
		logging.DebugLog(err)
		return false
	}
	return true
}

//添加用户收藏的帖子
//insert into `topiclikes` values(xxx)
func CreateTopicLikes(likes *TopicLikes) bool {
	if models.DB.NewRecord(likes) {
		models.DB.Debug().Create(likes)
		return !models.DB.NewRecord(likes)
	}
	return false
}
