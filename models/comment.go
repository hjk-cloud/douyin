package models

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/utils"
	"gorm.io/gorm"
	"sync"
)

type Comment struct {
	Id         int    `json:"id,omitempty"`
	userId     int    `json:"user_id"`
	User       User   `json:"user"`
	videoId    int    `json:"video_id"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

func (Comment) TableName() string {
	return "comment"
}

type CommentDao struct {
}

var commentDao *CommentDao
var commentOnce sync.Once

func NewCommentDaoInstance() *CommentDao {
	commentOnce.Do(
		func() {
			commentDao = &CommentDao{}
		})
	return commentDao
}

func (*CommentDao) MQueryCommentById(videoId int) []Comment {
	var comments []Comment
	err := db.Where("video_id = ?", videoId).Find(&comments).Error
	fmt.Println("comments----------", comments)
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	if err != nil {
		utils.Logger.Error("find comments error:" + err.Error())
		return nil
	}
	return comments
}
