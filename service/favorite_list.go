package service

import (
	"github.com/hjk-cloud/douyin/models"
	"github.com/hjk-cloud/douyin/util/jwt"
)

type FavoriteListFlow struct {
	Token  string
	UserId int
	Videos []*models.Video
}

func FavoriteList(token string, userId int) ([]*models.Video, error) {
	return NewFavoriteListFlow(token, userId).Do()
}

func NewFavoriteListFlow(token string, userId int) *FavoriteListFlow {
	return &FavoriteListFlow{
		Token:  token,
		UserId: userId,
	}
}

func (f *FavoriteListFlow) Do() ([]*models.Video, error) {
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

func (f *FavoriteListFlow) checkParam() error {
	//fmt.Println("favoriteService---Token----", f.Token)
	//fmt.Println("favoriteService---UserId----", f.UserId)
	if _, err := jwt.JWTAuth(f.Token); err != nil {
		return err
	}
	return nil
}

func (f *FavoriteListFlow) prepareData() error {

	return nil
}

func (f *FavoriteListFlow) packData() error {
	favoriteDao := models.NewFavoriteDaoInstance()
	videoDao := models.NewVideoDaoInstance()
	relationDao := models.NewRelationDaoInstance()

	videoIds := favoriteDao.QueryFavoriteVideo(f.UserId)
	f.Videos = videoDao.MQueryVideoByIds(videoIds)

	for i := range f.Videos {
		videoDao.BuildAuthor(f.Videos[i])
		f.Videos[i].Author.IsFollow = relationDao.QueryRelationState(f.UserId, f.Videos[i].AuthorId)
		f.Videos[i].IsFavorite = favoriteDao.QueryFavoriteState(f.UserId, f.Videos[i].Id)
		//fmt.Println("service----Videos[i]---", f.Videos[i])
	}

	return nil
}
