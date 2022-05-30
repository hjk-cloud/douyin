package models

import (
	"github.com/hjk-cloud/douyin/util"
	"gorm.io/gorm"
	"sync"
)

type Video struct {
	Id            int    `gorm:"column:id;type:int" json:"id,omitempty"`
	AuthorId      int    `gorm:"column:author_id" json:"author_id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `gorm:"column:play_url;type:varchar(255)" json:"play_url"`
	CoverUrl      string `gorm:"column:cover_url;type:varchar(255)" json:"cover_url"`
	FavoriteCount int    `json:"favorite_count,omitempty"`
	CommentCount  int    `json:"comment_count,omitempty"`
	Title         string `json:"title,omitempty"`
	IsFavorite    bool   `json:"is_favorite"`
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

func (*VideoDao) BuildAuthor(video *Video) error {
	user, _ := NewUserDaoInstance().QueryUserById(video.AuthorId)
	video.Author = *user
	return nil
}

func (*VideoDao) MQueryVideo(videos *[]*Video) error {
	err := db.Order("id desc").Limit(30).Find(&videos).Error
	if err == gorm.ErrRecordNotFound {
		return err
	}
	if err != nil {
		return err
	}
	return nil
}

func (*VideoDao) MQueryVideoByToken(token string, videos []*Video) error {
	//if err == gorm.ErrRecordNotFound {
	//	return err
	//}
	//err = db.Where("author_id = ?", user.Id).Find(&videos).Error
	//if err == gorm.ErrRecordNotFound {
	//	return err
	//}
	//if err != nil {
	//	util.Logger.Error("find videos by token error:" + err.Error())
	//	return err
	//}
	//for i := range videos {
	//	NewVideoDaoInstance().BuildAuthor(videos[i])
	//}
	return nil
}

func (*VideoDao) MQueryVideoByIds(videoIds []int) []*Video {
	var videos []*Video
	err := db.Where("id in ?", videoIds).Find(&videos).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	if err != nil {
		return nil
	}
	for i := range videos {
		NewVideoDaoInstance().BuildAuthor(videos[i])
	}
	return videos
}

//王硕-------------------通过视频id查找并返回对应的所有视频
func (*VideoDao) MQueryVideoByAuthorIds(videoIds []int) []Video {
	var videos []Video
	err := db.Where("id in ?", videoIds).Find(&videos).Error

	if err == gorm.ErrRecordNotFound {
		return nil
	}
	if err != nil {
		util.Logger.Error("find videos by ids error:" + err.Error())
		return nil
	}
	return videos
}

//王硕------------------------通过用户id查找该id下发布的所有视频的id
func (*VideoDao) QueryPublishVideoList(UserId int) []int {
	ids := make([]int, 0)
	err := db.Table("video").Select("author_id").Where("id = ?", UserId).Find(&ids).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	if err != nil {
		util.Logger.Error("find videoIds by author_id error:" + err.Error())
		return nil
	}
	return ids
}

func (*VideoDao) PublishVideo(video *Video) error {
	err := db.Select("author_id", "play_url", "cover_url", "title", "created_at").
		Create(&video).Error
	if err != nil {
		return err
	}
	return nil
}
