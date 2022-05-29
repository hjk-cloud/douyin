package service

import (
	"errors"
	"fmt"
	"github.com/hjk-cloud/douyin/models"
	"github.com/hjk-cloud/douyin/util/jwt"
)

func RelationAction(token string, userId int, toUserId int, actionType string) error {
	return NewRelationActionFlow(token, userId, toUserId, actionType).Do()
}

func NewRelationActionFlow(token string, userId int, toUserId int, actionType string) *RelationActionFlow {
	return &RelationActionFlow{
		Token:      token,
		userId:     userId,
		toUserId:   toUserId,
		actionType: actionType,
	}
}

type RelationActionFlow struct {
	Token      string
	userId     int
	toUserId   int
	actionType string
}

func (f *RelationActionFlow) Do() error {
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

func (f *RelationActionFlow) checkParam() error {
	fmt.Println("service-------", f.Token)
	fmt.Println("service-------", f.userId)
	fmt.Println("service-------", f.toUserId)
	fmt.Println("service-------", f.actionType)
	userId, err := jwt.JWTAuth(f.Token)
	fmt.Println("service---parseToken---", userId)
	if err != nil {
		return err
	}
	f.userId = userId
	return nil
}

func (f *RelationActionFlow) prepareData() error {
	relationDao := models.NewRelationDaoInstance()
	if f.actionType == "1" {
		relation := models.Relation{
			UserId:   f.userId,
			ToUserId: f.toUserId,
		}
		if err := relationDao.CreateRelation(relation); err != nil {
			return err
		}
	} else if f.actionType == "2" {
		relation, err := relationDao.QueryRelation(f.userId, f.toUserId)
		if err != nil {
			return errors.New("未关注")
		}
		if err := relationDao.DeleteRelation(relation); err != nil {
			return err
		}
	}
	return nil
}

func (f *RelationActionFlow) packData() error {

	return nil
}
