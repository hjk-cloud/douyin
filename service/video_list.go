package service

import (
	"github.com/hjk-cloud/douyin/models"
	"time"
)

type VideoListFlow struct {
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

//TODO 通过token获取userId，查找userId和video列表中每个video的点赞关系
func (f *VideoListFlow) prepareData() error {

	return nil
}

func (f *VideoListFlow) packData() error {
	videoDao := models.NewVideoDaoInstance()
	videoDao.MQueryVideo(&f.Videos)
	for i := range f.Videos {
		videoDao.BuildAuthor(f.Videos[i])
	}
	return nil
}
