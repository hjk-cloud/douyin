package controller

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	user, err := model.NewUserDaoInstance().QueryUserByToken(token)
	if err == nil {
		c.JSON(http.StatusOK, model.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
	//userIdString := c.Query("user_id")
	//userId, _ := strconv.Atoi(userIdString)
	userId := user.Id
	videoIdString := c.Query("video_id")
	actionType := c.Query("action_type") //1-点赞，2-取消点赞
	videoId, _ := strconv.Atoi(videoIdString)
	var favorite model.Favorite
	favorite.UserId = userId
	favorite.VideoId = videoId
	if actionType == "1" {
		if err := model.NewFavoriteDaoInstance().CreateFavorite(favorite); err == nil {
			c.JSON(http.StatusOK, model.Response{StatusCode: 0})
		}
	} else if actionType == "2" {
		if err := model.NewFavoriteDaoInstance().DeleteFavorite(favorite); err == nil {
			c.JSON(http.StatusOK, model.Response{StatusCode: 0})
		} else {
			c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "operate fail..."})
		}
	}
}

func FavoriteList(c *gin.Context) {
	//token := c.Query("token")
	userIdString := c.Query("user_id")
	userId, _ := strconv.Atoi(userIdString)
	videoIds := model.NewFavoriteDaoInstance().QueryFavoriteVideo(userId)
	videos := model.NewVideoDaoInstance().MQueryVideoByIds(videoIds)
	c.JSON(http.StatusOK, VideoListResponse{
		Response: model.Response{
			StatusCode: 0,
		},
		VideoList: videos,
	})
}
