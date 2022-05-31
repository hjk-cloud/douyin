package service

import (
	"github.com/hjk-cloud/douyin/models"
	"github.com/hjk-cloud/douyin/util/jwt"
	"time"
)

type VideoListFlow struct {
	UserId     int
	Token      string
	LatestTime time.Time
	Videos     []*models.Video
	NextTime   int64
}

func VideoListWithToken(token string, latestTime time.Time) ([]*models.Video, error) {
	return NewVideoListWithTokenFlow(token, latestTime).Do()
}

func NewVideoListWithTokenFlow(token string, latestTime time.Time) *VideoListFlow {
	return &VideoListFlow{Token: token, LatestTime: latestTime}
}

func (f *VideoListFlow) Do() ([]*models.Video, error) {
	if f.Token != "-" {
		if err := f.prepareData(); err != nil {
			return nil, err
		}
	}
	if err := f.packData(); err != nil {
		return nil, err
	}
	return f.Videos, nil
}

func (f *VideoListFlow) prepareData() error {
	userId, err := jwt.JWTAuth(f.Token)
	if err != nil {
		return err
	}
	f.UserId = userId
	return nil
}

func (f *VideoListFlow) packData() error {
	videoDao := models.NewVideoDaoInstance()
	relationDao := models.NewRelationDaoInstance()
	favoriteDao := models.NewFavoriteDaoInstance()
	commentDao := models.NewCommentDaoInstance()

	videoDao.MQueryVideo(&f.Videos)
	for i := range f.Videos {
		videoDao.BuildAuthor(f.Videos[i])

		user, _ := UserInfo(f.Token, f.Videos[i].AuthorId)
		f.Videos[i].Author = *user
		//fmt.Println(f.Videos[i].Author)
		f.Videos[i].Author.IsFollow = relationDao.QueryRelationState(f.UserId, f.Videos[i].AuthorId)
		f.Videos[i].IsFavorite = favoriteDao.QueryFavoriteState(f.UserId, f.Videos[i].Id)
		//fmt.Println("service--",f.UserId,"----",f.Videos[i].Id,"--isFavorite---", f.Videos[i].IsFavorite)
		f.Videos[i].FavoriteCount = favoriteDao.QueryVideoFavoriteCount(f.Videos[i].Id)
		f.Videos[i].CommentCount = commentDao.QueryCommentCount(f.Videos[i].Id)
	}
	return nil
}
