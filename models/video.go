package models

import (
	"github.com/RaymondCode/simple-demo/utils"
	"gorm.io/gorm"
	"sync"
)

type Video struct {
	Id            int64  `gorm:"column:id;type:int" json:"id,omitempty"`
	AuthorId      int64  `gorm:"column:author_id" json:"author_id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `gorm:"column:play_url;type:varchar(255)" json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `gorm:"column:play_url;type:varchar(255)" json:"cover_url,omitempty"`
	FavoriteCount int64  `gorm:"column:favorite_count;type:int" json:"favorite_count,omitempty"`
	CommentCount  int64  `gorm:"column:comment_count;type:int" json:"comment_count,omitempty"`
	IsFavorite    bool   `gorm:"column:is_favorite;type:tinyint(1)" json:"is_favorite,omitempty"`
	Title         string `json:"title,omitempty"`
}

func (Video) TableName() string {
	return "video"
}

type VideoDao struct {
}

var videoDao *VideoDao
var videoOnce sync.Once

func NewVideoDaoInstance() *VideoDao {
	videoOnce.Do(
		func() {
			videoDao = &VideoDao{}
		})
	return videoDao
}

func (*VideoDao) BuildAuthor(video Video) User {
	user, _ := NewUserDaoInstance().QueryUserById(video.AuthorId)
	video.Author = *user
	return video.Author
}

func (*VideoDao) MQueryVideo() []Video {
	var videos []Video
	err := db.Limit(30).Find(&videos).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	if err != nil {
		utils.Logger.Error("find videos error:" + err.Error())
		return nil
	}
	for i := range videos {
		videos[i].Author = NewVideoDaoInstance().BuildAuthor(videos[i])
	}
	return videos
}

func (*VideoDao) PublishVideo(video *Video) error {
	err := db.Create(&video).Error
	if err != nil {
		utils.Logger.Error("create video err:" + err.Error())
		return err
	}
	return nil
}
