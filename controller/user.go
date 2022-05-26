package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]models.User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

var userIdSequence = int(1)

type UserLoginResponse struct {
	models.Response
	UserId int    `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	models.Response
	User models.User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password
	_, err := models.NewUserDaoInstance().QueryUserByName(username)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: models.Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		newUser := models.User{
			Name:     username,
			Password: password,
			Token:    username + password,
		}
		err = models.NewUserDaoInstance().Register(&newUser)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: models.Response{StatusCode: 0},
			UserId:   userIdSequence,
			Token:    token,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	user, err := models.NewUserDaoInstance().Login(username, password)
	fmt.Println("login-----user_id----", user.Id)
	if err == nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: models.Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: models.Response{StatusCode: 1, StatusMsg: "login----User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {
	userIdString := c.Query("user_id")
	userId, _ := strconv.Atoi(userIdString)
	fmt.Println("user_id----------", userId)
	token := c.Query("token")
	user, err := models.NewUserDaoInstance().QueryUserByToken(token)
	if user != nil && err == nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: models.Response{StatusCode: 0},
			User:     *user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: models.Response{StatusCode: 1, StatusMsg: "user info---User doesn't exist"},
		})
	}
}
