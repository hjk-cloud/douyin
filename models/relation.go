package models

import (
	"github.com/RaymondCode/simple-demo/util"
	"gorm.io/gorm"
	"sync"
)

type Relation struct {
	UserId   int `json:"user_id"`
	ToUserId int `json:"to_user_id"`
}

func (Relation) TableName() string {
	return "relation"
}

type RelationDao struct {
}

var relationDao *RelationDao
var relationOnce sync.Once

func NewRelationDaoInstance() *RelationDao {
	relationOnce.Do(
		func() {
			relationDao = &RelationDao{}
		})
	return relationDao
}

func (*RelationDao) CreateRelation(relation Relation) error {
	if err := db.Create(&relation).Error; err != nil {
		util.Logger.Error("insert relation err:" + err.Error())
		return err
	}
	return nil
}

func (*RelationDao) DeleteRelation(relation Relation) error {
	if err := db.Where("user_id = ? and to_user_id = ?", relation.UserId, relation.ToUserId).Delete(&relation).Error; err != nil {
		util.Logger.Error("delete relation err:" + err.Error())
		return err
	}
	return nil
}

func (*RelationDao) QueryRelationByUserId(userId int) []int {
	ids := make([]int, 0)
	err := db.Table("relation").Select("to_user_id").Where("user_id = ?", userId).Find(&ids).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	if err != nil {
		util.Logger.Error("find relations by user_id error:" + err.Error())
		return nil
	}
	return ids
}

func (*RelationDao) QueryRelationByToUserId(userId int) []int {
	ids := make([]int, 0)
	err := db.Table("relation").Select("user_id").Where("to_user_id = ?", userId).Find(&ids).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	if err != nil {
		util.Logger.Error("find relations by user_id error:" + err.Error())
		return nil
	}
	return ids
}

func (*RelationDao) QueryRelation(userId int, toUserId int) (Relation, error) {
	var relation Relation
	err := db.Where("user_id = ? and to_user_id = ?", userId, toUserId).Find(&relation).Error
	if err == gorm.ErrRecordNotFound {
		return relation, err
	}
	if err != nil {
		util.Logger.Error("find relations by user_id error:" + err.Error())
		return relation, err
	}
	return relation, nil
}
