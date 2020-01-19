package Relation

import (
	"w3fy/models"
	"w3fy/pkg/logging"
)

type Relation struct {
	models.Model

	FromId int `json:"from_id"`
	ToId   int `json:"to_id"`
}

func (Relation) TableName() string {
	return "relation"
}

//获取某用户关注的用户
//select `to_id` from Relation where(`from_id`=xx)
func GetFocus(pageNum int, pageSize int, fid int) (relation []Relation, err error) {
	if err = models.DB.Debug().Where("from_id=?", fid).Offset(pageNum).Limit(pageSize).Find(&relation).Error; err != nil {
		logging.DebugLog(err)
	}
	return
}

//获取某用户的粉丝
//select `from_id` from Relation where(`to_id`=xx)
func GerFans(pageNum int, pageSize int, tid int) (relation []Relation, err error) {
	if err = models.DB.Debug().Where("to_id=?", tid).Offset(pageNum).Limit(pageSize).Find(&relation).Error; err != nil {
		logging.DebugLog(err)
	}
	return
}

//添加一条relation数据
//insert into `Relation` values(xxx)
func CreateRelation(relation Relation) bool {
	if models.DB.NewRecord(relation) {
		models.DB.Debug().Create(relation)
		return !models.DB.NewRecord(relation)
	}
	return false
}

//删除一条relation数据
//delete from `Relation`  where (`from_id`==xx and `to_id`==xx)
func DeleteRelation(relation *Relation) bool {
	if err := models.DB.Debug().Unscoped().Delete(relation).Error; err != nil {
		logging.DebugLog(err)
		return false
	}
	return true
}
