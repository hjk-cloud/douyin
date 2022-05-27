package models

import (
	"github.com/hjk-cloud/douyin/util"
	"gorm.io/gorm"
	"sync"
)

type Favorite struct {
	UserId  int `json:"user_id"`
	VideoId int `json:"video_id"`
}

func (Favorite) tableName() string {
	return "favorite"
}

type FavoriteDao struct {
}

var favoriteDao *FavoriteDao
var favoriteOnce sync.Once

func NewFavoriteDaoInstance() *FavoriteDao {
	favoriteOnce.Do(
		func() {
			favoriteDao = &FavoriteDao{}
		})
	return favoriteDao
}

func (*FavoriteDao) CreateFavorite(favorite Favorite) error {
	if err := db.Table("favorite").Create(&favorite).Error; err != nil {
		util.Logger.Error("insert favorite err:" + err.Error())
		return err
	}
	return nil
}

func (*FavoriteDao) DeleteFavorite(favorite Favorite) error {
	if err := db.Table("favorite").Where("user_id = ? and video_id = ?", favorite.UserId, favorite.VideoId).Delete(&favorite).Error; err != nil {
		util.Logger.Error("delete favorite err:" + err.Error())
		return err
	}
	return nil
}

func (*FavoriteDao) QueryFavoriteCount(videoId int) (int, error) {
	var count int64
	err := db.Table("favorite").Where("video_id = ?", videoId).Count(&count).Error
	if err == gorm.ErrRecordNotFound {
		return 0, err
	}
	if err != nil {
		util.Logger.Error("find relations by user_id error:" + err.Error())
		return 0, err
	}
	return int(count), nil
}

func (*FavoriteDao) QueryFavoriteVideo(userId int) []int {
	videoIds := make([]int, 0)
	err := db.Table("favorite").Select("video_id").Where("user_id = ?", userId).Find(&videoIds).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	if err != nil {
		util.Logger.Error("find favorite by user_id error:" + err.Error())
		return nil
	}
	return videoIds
}
