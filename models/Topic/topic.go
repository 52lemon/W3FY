package Topic

import (
	"w3fy/models"
	"w3fy/pkg/logging"
)

type Topic struct {
	models.Model

	Tag       string `json:"tag"`
	Uid       int    `json:"uid"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	IsDeleted int    `json:"-"`
}

//用户创建帖子 成功条件:(1)该帖子为新记录
// insert into `topic` values(xxx,xxx,xxx,xxx)
func CreateTopic(topic *Topic) bool {
	if models.DB.NewRecord(topic) {
		models.DB.Debug().Create(topic)
		return !models.DB.NewRecord(topic)
	}
	return false
}

//用户修改帖子 成功条件:(1)该帖子存在
// update `topic` set aa=xx,bb=xx
func UpdateTopic(topic *Topic, data interface{}) bool {
	if err := models.DB.Where(topic).Updates(data).Error; err == nil {
		return true
	}
	return false
}

//用户删除帖子 成功条件:(1)该帖子存在.(2)该帖子不是在逻辑删除状态
// select `is_deleted` from topic (if is_deleted==0) update  `topic` set `is_deleted`=0 where `id` =xx
func DeleteTopic(id int) bool {
	var find Topic
	if err := models.DB.Model(find).Where("id=?", id).First(&find).Error; err == nil && find.IsDeleted == 0 {
		if err = models.DB.Debug().Model(&find).Update("is_deleted", 1).Error; err == nil {
			return true
		}
	}
	return false
}

//获取用户个人帖子数据 成功条件:(1)帖子为逻辑删除否
//select * from `topic` where `uid`=xxx and `is_deleted` =0
func GetUserTopic(uid int) (topics []Topic) {
	if err := models.DB.Debug().Model(Topic{}).Where(map[string]interface{}{"uid": uid, "is_deleted": 0}).Find(&topics).Error; err != nil {
		logging.DebugLog(err)
	}
	return
}

//获取用户帖子，按更新时间升序 成功条件:(1)帖子为逻辑删除否
//select * from `topic` where `is_deleted` = 0 order by `updated_at` desc
func GetTopic() (topics []Topic) {
	if err := models.DB.Debug().Model(Topic{}).Where("is_deleted=?", 0).Order("updated_at desc").Find(&topics).Error; err != nil {
		logging.DebugLog(err)
	}
	return
}
