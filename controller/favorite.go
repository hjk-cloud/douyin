package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hjk-cloud/douyin/models"
	"github.com/hjk-cloud/douyin/service"
	"net/http"
	"strconv"
)

type FavoriteVideoListResponse struct {
	Response
	VideoList []*models.Video `json:"video_list"`
}

func FavoriteAction(c *gin.Context) {
	//token := c.Query("token")
	//user, err := models.NewUserDaoInstance().QueryUserByToken(token)
	//if err == nil {
	//	c.JSON(http.StatusOK, Response{StatusCode: 0})
	//} else {
	//	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	//}
	////userIdString := c.Query("user_id")
	////userId, _ := strconv.Atoi(userIdString)
	//userId := user.Id
	//videoIdString := c.Query("video_id")
	//actionType := c.Query("action_type") //1-点赞，2-取消点赞
	//videoId, _ := strconv.Atoi(videoIdString)
	//var favorite models.Favorite
	//favorite.UserId = userId
	//favorite.VideoId = videoId
	//if actionType == "1" {
	//	if err := models.NewFavoriteDaoInstance().CreateFavorite(favorite); err == nil {
	//		c.JSON(http.StatusOK, Response{StatusCode: 0})
	//	}
	//} else if actionType == "2" {
	//	if err := models.NewFavoriteDaoInstance().DeleteFavorite(favorite); err == nil {
	//		c.JSON(http.StatusOK, Response{StatusCode: 0})
	//	} else {
	//		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "operate fail..."})
	//	}
	//}
}

func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	userIdString := c.Query("user_id")
	userId, _ := strconv.Atoi(userIdString)

	videos, _ := service.FavoriteList(token, userId)

	c.JSON(http.StatusOK, FavoriteVideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: videos,
	})
}
