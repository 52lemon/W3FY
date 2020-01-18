package Comments

import (
	"w3fy/models"
	"w3fy/pkg/logging"
)

type Comment struct {
	models.Model

	TopId    int    `json:"top_id"`
	FromId   int    `json:"from_id"`
	ToId     int    `json:"to_id"`
	Comments string `json:"comments"`
}

func (Comment) TableName() string {
	return "comment"
}

//添加评论 成功条件:(1)帖子存在
//insert into `comment`(`top_id`,`father_id`,`from_id`,`to_id`,`comments`,`created_at`)values(xxx,xxx,xxx,xxx,xxx)
func CreateComment(comment Comment) bool {
	if models.DB.NewRecord(comment) {
		models.DB.Debug().Create(&comment)
		return !models.DB.NewRecord(comment)
	}
	return false
}

//获取某帖子的评论 成功条件(1)帖子存在
//select * from `comment` where(`top_ic`=xx)
func GetCommentByTid(pageNum int, pageSize int, tid int) (comments []Comment, err error) {
	if err = models.DB.Debug().Model(&Comment{}).Where("top_id=?", tid).Offset(pageNum).Limit(pageSize).Find(&comments).Error; err != nil {
		logging.DebugLog(err)
	}
	return
}

//获取评论人id
//select `from_id` from `comment` where (`id`=xx)
func GetFromId(id int) int {
	var find Comment
	if err := models.DB.Debug().Where("id=?", id).First(&find); err != nil {
		logging.DebugLog(err)
	}
	return find.FromId
}

//逻辑删除评论  成功条件:(1)删除人id=from_id (2)帖子的is_deleted=0
// update `comment` set `is_deleted` =1,`comments`="评论已被删除" where(`from_id`=id and select `is_deleted` from `topic` where(`id`=tid) = 0)
func DeleteComment(id int) bool {
	if err := models.DB.Debug().Model(&Comment{}).Where("id=?", id).Updates(map[string]interface{}{"comments": "评论已被删除"}); err != nil {
		logging.DebugLog(err)
		return false
	}
	return true
}
