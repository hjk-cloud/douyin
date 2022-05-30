package service

import (
	"github.com/hjk-cloud/douyin/models"
	"github.com/hjk-cloud/douyin/util/jwt"
)

type PublishListFlow struct {
	Token  string
	UserId int
	Videos []models.Video
}

func PublishList(token string, userId int) ([]models.Video, error) {
	return NewPublishListWithTokenFlow(token, userId).Do()
}

func NewPublishListWithTokenFlow(token string, userId int) *PublishListFlow {
	return &PublishListFlow{
		Token:  token,
		UserId: userId,
	}
}

func (f *PublishListFlow) Do() ([]models.Video, error) {
	if err := f.checkParam(); err != nil {
		return nil, err
	}
	if err := f.prepareData(); err != nil {
		return nil, err
	}
	if err := f.packData(); err != nil {
		return nil, err
	}
	return f.Videos, nil
}

func (f *PublishListFlow) checkParam() error {
	//fmt.Println("service---Token----", f.Token)
	//fmt.Println("service---AuthorId----", f.UserId)
	_, err := jwt.JWTAuth(f.Token)
	if err != nil {
		return err
	}
	return nil
}

func (f *PublishListFlow) prepareData() error {
	userId, err := jwt.JWTAuth(f.Token)
	if err != nil {
		return err
	}
	f.UserId = userId
	return nil
}

func (f *PublishListFlow) packData() error {
	videoDao := models.NewVideoDaoInstance()
	favoriteDao := models.NewFavoriteDaoInstance()

	videos := videoDao.QueryPublishVideoList(f.UserId)
	//fmt.Println("videoId-----------", videos)
	f.Videos = videoDao.MQueryVideoByAuthorIds(videos)
	for i := range f.Videos {
		f.Videos[i].IsFavorite = favoriteDao.QueryFavoriteState(f.UserId, f.Videos[i].Id)
		f.Videos[i].FavoriteCount = favoriteDao.QueryVideoFavoriteCount(f.Videos[i].Id)
	}
	return nil
}
