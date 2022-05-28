package models

import (
	"errors"
	"github.com/hjk-cloud/douyin/util"
	"gorm.io/gorm"
	"sync"
)

type User struct {
	Id            int    `gorm:"column:id;type:int" json:"id,omitempty"`
	Name          string `gorm:"column:name;type:varchar(10);" json:"name,omitempty"`
	Password      string `gorm:"column:password;type:varchar(10);" json:"password"`
	FollowCount   int    `gorm:"column:follow_count;type:int" json:"follow_count,omitempty"`
	FollowerCount int    `gorm:"column:follower_count;type:int" json:"follower_count,omitempty"`
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
	err := db.Omit("token").Create(&user).Error
	if err != nil {
		return errors.New("创建用户失败")
	}
	return nil
}

func (*UserDao) QueryUserById(id int) (*User, error) {
	var user User
	err := db.Where("id = ?", id).Take(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, errors.New("未查询用户id")
	}
	return &user, nil
}

func (*UserDao) MQueryUserById(ids []int) []User {
	var users []User
	err := db.Where("id in ?", ids).Find(&users).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	if err != nil {
		return nil
	}
	return users
}

func (*UserDao) QueryUserByToken(token string) (*User, error) {
	var user User
	err := db.Where("token = ?", token).Take(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		util.Logger.Error("find user by  token err:" + err.Error())
		return nil, err
	}
	return &user, nil
}

func (*UserDao) QueryUserByName(name string) (int, error) {
	var count int64
	err := db.Where("name = ?", name).Count(&count).Error
	if err == gorm.ErrRecordNotFound {
		return 1, nil
	}
	return int(count), nil
}

func (*UserDao) Login(name string, password string) (int, error) {
	var user User
	err := db.Where("name = ? AND password = ?", name, password).Take(&user).Error
	if err == gorm.ErrRecordNotFound {
		return 0, err
	}
	if err != nil {
		util.Logger.Error("login err:" + err.Error())
		return 0, err
	}
	return user.Id, nil
}
