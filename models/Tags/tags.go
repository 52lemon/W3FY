package Tags

import (
	"w3fy/models"
	"w3fy/pkg/logging"
)

type Tags struct {
	models.Model
	Name string `json:"name"`
}

func (Tags) TableName() string {
	return "tags"
}

//添加节点
//insert into tags (`name`)values(`xxx`)
func CreateTag(tags *Tags) bool {
	if models.DB.NewRecord(tags) {
		models.DB.Debug().Create(tags)
		return !models.DB.NewRecord(tags)
	}
	return false
}

//获取节点列表
//select `name` from tags
func GetTags() (tags []Tags) {
	if err := models.DB.Debug().Model(&Tags{}).Select("name").Find(&tags).Error; err != nil {
		logging.DebugLog(err)
	}
	return
}

//删除节点
//delete from `tags`  where(`name`=xxx)
func DeleteTag(name string) bool {
	if err := models.DB.Debug().Unscoped().Where("name=?", name).Delete(&Tags{}).Error; err != nil {
		logging.DebugLog(err)
		return false
	}
	return true
}
