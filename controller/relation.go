package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hjk-cloud/douyin/models"
	"net/http"
	"strconv"
)

type UserListResponse struct {
	Response
	UserList []models.User `json:"user_list"`
}

func RelationAction(c *gin.Context) {
	token := c.Query("token")
	//从前端获取到的user_id一直为0，目前解决方法是获取token，再根据token获取当前用户user_id
	//userIdString := c.Query("user_id")
	//userId, _ := strconv.Atoi(userIdString)
	toUserIdString := c.Query("to_user_id")
	toUserId, _ := strconv.Atoi(toUserIdString)
	actionType := c.Query("action_type") //1-关注，2-取消关注
	user, err := models.NewUserDaoInstance().QueryUserByToken(token)
	if err == nil {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	fmt.Println("user_id----------", user.Id)
	fmt.Println("to_user_id----------", toUserId)
	fmt.Println("action_type----------", actionType)
	if actionType == "1" {
		relation, _ := models.NewRelationDaoInstance().QueryRelation(user.Id, toUserId)
		relation.UserId = user.Id
		relation.ToUserId = toUserId
		models.NewRelationDaoInstance().CreateRelation(relation)
	} else {
		relation, err := models.NewRelationDaoInstance().QueryRelation(user.Id, toUserId)
		if err != nil {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "no Relation !!!!!!"})
			return
		}
		models.NewRelationDaoInstance().DeleteRelation(relation)
	}
}

//关注列表
func FollowList(c *gin.Context) {
	userIdString := c.Query("user_id")
	userId, _ := strconv.Atoi(userIdString)
	ids := models.NewRelationDaoInstance().QueryRelationByUserId(userId)
	users := models.NewUserDaoInstance().MQueryUserById(ids)
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: users,
	})
}

//粉丝列表
func FollowerList(c *gin.Context) {
	userIdString := c.Query("user_id")
	userId, _ := strconv.Atoi(userIdString)
	ids := models.NewRelationDaoInstance().QueryRelationByToUserId(userId)
	users := models.NewUserDaoInstance().MQueryUserById(ids)
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: users,
	})
}
