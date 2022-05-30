package service

import "github.com/hjk-cloud/douyin/models"

type FavoriteActionFlow struct {
	UserId     int
	Token      string
	VideoId    int
	ActionType string
}

func FavoriteAction(userId int, token string, videoId int, actionType string) error {
	return NewFavoriteActionFlow(userId, token, videoId, actionType).Do()
}

func NewFavoriteActionFlow(userId int, token string, videoId int, actionType string) *FavoriteActionFlow {
	return &FavoriteActionFlow{
		UserId:     userId,
		Token:      token,
		VideoId:    videoId,
		ActionType: actionType,
	}
}

func (f *FavoriteActionFlow) Do() error {
	if err := f.checkParam(); err != nil {
		return err
	}
	if err := f.prepareData(); err != nil {
		return err
	}
	if err := f.packData(); err != nil {
		return err
	}
	return nil
}

func (f *FavoriteActionFlow) checkParam() error {

	return nil
}

func (f *FavoriteActionFlow) prepareData() error {
	favoriteDao := models.NewFavoriteDaoInstance()
	favorite := models.Favorite{
		UserId:  f.UserId,
		VideoId: f.VideoId,
	}
	var err error

	if f.ActionType == "1" {
		err = favoriteDao.CreateFavorite(favorite)
	} else {
		err = favoriteDao.DeleteFavorite(favorite)
	}
	if err != nil {
		return err
	}
	return nil
}

func (f *FavoriteActionFlow) packData() error {

	return nil
}
