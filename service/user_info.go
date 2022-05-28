package service

import (
	"errors"
	"github.com/hjk-cloud/douyin/models"
	"github.com/hjk-cloud/douyin/util/jwt"
	"strconv"
)

type UserInfoFlow struct {
	UserId        string
	Token         string
	User          *models.User
	FollowCount   int
	FollowerCount int
}

func UserInfo(token string, userId string) (*models.User, error) {
	return NewUserInfoFlow(token, userId).Do()
}

func NewUserInfoFlow(token string, userId string) *UserInfoFlow {
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
	if f.UserId == "" {
		return errors.New("id为空")
	}
	return nil
}

func (f *UserInfoFlow) prepareData() error {
	userId, err := jwt.JWTAuth(f.Token)
	if err != nil {
		return err
	}
	relationDao := models.NewRelationDaoInstance()
	followCount, err := relationDao.QueryRelationCountByUserId(userId)
	if err != nil {
		return err
	}
	f.FollowCount = followCount
	followerCount, err := relationDao.QueryRelationCountByToUserId(userId)
	if err != nil {
		return err
	}
	f.FollowerCount = followerCount
	return nil
}

func (f *UserInfoFlow) packData() error {
	userDao := models.NewUserDaoInstance()
	userId, _ := strconv.Atoi(f.UserId)
	user, err := userDao.QueryUserById(userId)
	if err != nil {
		return err
	}
	f.User = user
	f.User.FollowCount = f.FollowCount
	f.User.FollowerCount = f.FollowerCount
	return nil
}
