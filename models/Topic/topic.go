package Topic

import (
	"fmt"
	"w3fy/models"
	"w3fy/pkg/logging"
)

type Topic struct {
	models.Model

	Tag     string `json:"tag"`
	Uid     int    `json:"uid"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (Topic) TableName() string {
	return "topic"
}

//获取某一指定的帖子
//select * from `topic` where `id`=xx
func GetTopicById(id int) (topic Topic, err error) {
	if err = models.DB.Debug().Where("id=?", id).First(&topic).Error; err != nil {
		logging.DebugLog(err)
	}
	return
}

//获取帖子的uid
//select uid from `topic` where `id`=xx
func GetUidById(id int) (topic Topic) {
	if err := models.DB.Select("uid").Where("id=?", id).First(&topic).Error; err != nil {
		logging.DebugLog(err)
	}
	return
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
func UpdateTopic(id int, data interface{}) bool {
	if err := models.DB.Debug().Model(Topic{}).Where("id=?", id).Updates(data).Error; err == nil {
		return true
	}
	return false
}

//用户删除帖子 成功条件:(1)该帖子存在.(2)该帖子不是在逻辑删除状态
// select `is_deleted` from topic (if is_deleted==0) update  `topic` set `is_deleted`=0 where `id` =xx
func DeleteTopic(id int) bool {
	if err := models.DB.Debug().Where("id=?", id).Delete(&Topic{}).Error; err == nil {
		return true
	}
	return false
}

//获取用户个人帖子数据 成功条件:(1)帖子为逻辑删除否
//select * from `topic` where `uid`=xxx and `is_deleted` =0
func GetUserTopic(uid int) (topics []Topic, err error) {
	if err = models.DB.Debug().Model(Topic{}).Where(map[string]interface{}{"uid": uid}).Find(&topics).Error; err != nil {
		logging.DebugLog(err)
	}
	return
}

//分页获取用户帖子，按更新时间升序 成功条件:(1)帖子为逻辑删除否
func GetTopic(pageNum int, pageSize int) (topics []Topic, err error) {
	if err = models.DB.Debug().Model(&Topic{}).Offset(pageNum).Limit(pageSize).Order("updated_at desc").Find(&topics).Error; err != nil {
		logging.DebugLog(err)
	}
	return
}

//根据标签查询
func SearchInTag(pageNum int, pageSize int, data string) (topics []Topic, err error) {
	data = fmt.Sprintf("%%%s%%", data)
	if err = models.DB.Debug().Where("tag LIKE ?", data).Offset(pageNum).Limit(pageSize).Find(&topics).Error; err != nil {
		logging.DebugLog(err)
	}
	return
}

//根据标题查询
func SearchInTitle(pageNum int, pageSize int, data string) (topics []Topic, err error) {
	data = fmt.Sprintf("%%%s%%", data)
	if err = models.DB.Debug().Where("title LIKE ?", data).Offset(pageNum).Limit(pageSize).Find(&topics).Error; err != nil {
		logging.DebugLog(err)
	}
	return
}
