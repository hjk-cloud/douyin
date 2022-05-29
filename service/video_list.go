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

	videoDao.MQueryVideo(&f.Videos)
	for i := range f.Videos {
		videoDao.BuildAuthor(f.Videos[i])
		f.Videos[i].Author.IsFollow = relationDao.QueryRelation(f.UserId, f.Videos[i].AuthorId)
	}
	return nil
}
