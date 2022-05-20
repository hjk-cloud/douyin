package models

import "fmt"

type User struct {
	Id            int64  `gorm:"column:id;type:int" json:"id,omitempty"`
	Name          string `gorm:"column:name;type:varchar(10);" json:"name,omitempty"`
	FollowCount   int64  `gorm:"column:follow_count;type:int" json:"follow_count,omitempty"`
	FollowerCount int64  `gorm:"column:follower_count;type:int" json:"follower_count,omitempty"`
	IsFollow      bool   `gorm:"column:is_follow;type:tinyint(1)" json:"is_follow,omitempty"`
}

func (User) TableName() string {
	return "user"
}

func GetUserList() {
	data := make([]*User, 0)
	DB.Find(&data)
	for _, v := range data {
		fmt.Printf("user: %v \n", v)
	}
}
