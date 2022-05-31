package service

import (
	"errors"
	"fmt"
	"github.com/hjk-cloud/douyin/models"
	"github.com/hjk-cloud/douyin/util/jwt"
)

type UserInfoFlow struct {
	UserId         int
	Token          string
	User           *models.User
	FollowCount    int
	FollowerCount  int
	TotalFavorited int
	FavoriteCount  int
}

func UserInfo(token string, userId int) (*models.User, error) {
	return NewUserInfoFlow(token, userId).Do()
}

func NewUserInfoFlow(token string, userId int) *UserInfoFlow {
	return &UserInfoFlow{Token: token, UserId: userId}
}

func (f *UserInfoFlow) Do() (*models.User, error) {
	if err := f.checkParam(); err != nil {
		return nil, err
	}
	if err := f.prepareData(); err != nil {
		return nil, err
	}
	if err := f.packData(); err != nil {
		return nil, err
	}
	return f.User, nil
}

func (f *UserInfoFlow) checkParam() error {
	if f.UserId == 0 {
		return errors.New("id为空")
	}
	return nil
}

func (f *UserInfoFlow) prepareData() error {
	relationDao := models.NewRelationDaoInstance()
	favoriteDao := models.NewFavoriteDaoInstance()
	videoDao := models.NewVideoDaoInstance()
	var totalFavorited = 0

	userId, err := jwt.JWTAuth(f.Token)
	if err != nil {
		return err
	}
	//关注数
	followCount, err := relationDao.QueryRelationCountByUserId(userId)
	if err != nil {
		return err
	}
	f.FollowCount = followCount
	//粉丝数
	followerCount, err := relationDao.QueryRelationCountByToUserId(userId)
	if err != nil {
		return err
	}
	f.FollowerCount = followerCount
	//获赞数
	videoIds := videoDao.QueryPublishVideoList(userId)
	for i := range videoIds {
		totalFavorited += favoriteDao.QueryVideoFavoriteCount(videoIds[i])
	}
	f.TotalFavorited = totalFavorited
	//喜欢数
	f.FavoriteCount = favoriteDao.QueryUserFavoriteCount(f.UserId)
	fmt.Println("total---", f.TotalFavorited)
	fmt.Println("favoriteCount-------", f.FavoriteCount)
	return nil
}

func (f *UserInfoFlow) packData() error {
	userDao := models.NewUserDaoInstance()

	user, err := userDao.QueryUserById(f.UserId)
	if err != nil {
		return err
	}
	f.User = user
	f.User.FollowCount = f.FollowCount
	f.User.FollowerCount = f.FollowerCount
	f.User.TotalFavorited = f.TotalFavorited
	f.User.FavoriteCount = f.FavoriteCount
	return nil
}
