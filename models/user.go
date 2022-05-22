package models

import (
	"github.com/RaymondCode/simple-demo/utils"
	"gorm.io/gorm"
	"sync"
)

type User struct {
	Id            int64  `gorm:"column:id;type:int" json:"id,omitempty"`
	Name          string `gorm:"column:name;type:varchar(10);" json:"name,omitempty"`
	Password      string `gorm:"column:password;type:varchar(10);" json:"password"`
	FollowCount   int64  `gorm:"column:follow_count;type:int" json:"follow_count,omitempty"`
	FollowerCount int64  `gorm:"column:follower_count;type:int" json:"follower_count,omitempty"`
	IsFollow      bool   `gorm:"column:is_follow;type:tinyint(1)" json:"is_follow,omitempty"`
	Token         string `json:"token"`
}

func (User) TableName() string {
	return "user"
}

type UserDao struct {
}

var userDao *UserDao
var userOnce sync.Once

func NewUserDaoInstance() *UserDao {
	userOnce.Do(
		func() {
			userDao = &UserDao{}
		})
	return userDao
}

func (*UserDao) Register(user *User) error {
	err := db.Create(&user).Error
	if err != nil {
		utils.Logger.Error("create user err:" + err.Error())
		return err
	}
	return nil
}

func (*UserDao) QueryUserById(id int64) (*User, error) {
	var user User
	err := db.Where("id = ?", id).Take(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		utils.Logger.Error("find user by id err:" + err.Error())
		return nil, err
	}
	return &user, nil
}

func (*UserDao) QueryUserByToken(token string) (*User, error) {
	var user User
	err := db.Where("token = ?", token).Take(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		utils.Logger.Error("find user by  token err:" + err.Error())
		return nil, err
	}
	return &user, nil
}

func (*UserDao) QueryUserByName(name string) (*User, error) {
	var user User
	err := db.Where("name = ?", name).Take(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		utils.Logger.Error("find user by id err:" + err.Error())
		return nil, err
	}
	return &user, nil
}

func (*UserDao) Login(name string, password string) (*User, error) {
	var user User
	err := db.Where("name = ? AND password = ?", name, password).Take(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}
	if err != nil {
		utils.Logger.Error("login err:" + err.Error())
		return nil, err
	}
	return &user, nil
}
