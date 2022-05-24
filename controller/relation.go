package controller

import (
	"github.com/RaymondCode/simple-demo/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserListResponse struct {
	models.Response
	UserList []models.User `json:"user_list"`
}

//接到的user_id和to_user_id都是0，怀疑前端接口没写好
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	userIdString := c.Query("user_id")
	toUserIdString := c.Query("to_user_id")
	actionType := c.Query("action_type") //1-关注，2-取消关注
	userId, _ := strconv.Atoi(userIdString)
	toUserId, _ := strconv.Atoi(toUserIdString)
	if _, err := models.NewUserDaoInstance().QueryUserByToken(token); err == nil {
		c.JSON(http.StatusOK, models.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, models.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	relation, err := models.NewRelationDaoInstance().QueryRelation(userId, toUserId)
	if err != nil {
		c.JSON(http.StatusOK, models.Response{StatusCode: 1, StatusMsg: "Relation err!!!!!!"})
		return
	}
	if actionType == "1" {
		models.NewRelationDaoInstance().CreateRelation(relation)
	} else {
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
		Response: models.Response{
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
		Response: models.Response{
			StatusCode: 0,
		},
		UserList: users,
	})
}
