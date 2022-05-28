package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hjk-cloud/douyin/models"
	"net/http"
	"strconv"
)

func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	user, err := models.NewUserDaoInstance().QueryUserByToken(token)
	//模仿了一下修改，感觉没区别，不知道是不是改错了
	//username := c.Query("username")
	//password := c.Query("password")
	//
	//user, err := models.NewUserDaoInstance().QueryUserByToken(username + password)
	if err == nil {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
	//userIdString := c.Query("user_id")
	//userId, _ := strconv.Atoi(userIdString)
	userId := user.Id
	videoIdString := c.Query("video_id")
	actionType := c.Query("action_type") //1-点赞，2-取消点赞
	videoId, _ := strconv.Atoi(videoIdString)
	var favorite models.Favorite
	favorite.UserId = userId
	favorite.VideoId = videoId
	if actionType == "1" {
		if err := models.NewFavoriteDaoInstance().CreateFavorite(favorite); err == nil {
			c.JSON(http.StatusOK, Response{StatusCode: 0})
		}
	} else if actionType == "2" {
		if err := models.NewFavoriteDaoInstance().DeleteFavorite(favorite); err == nil {
			c.JSON(http.StatusOK, Response{StatusCode: 0})
		} else {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "operate fail..."})
		}
	}
	//todo favorite_count +1 / -1
	var video models.Video
	if actionType == "1" {
		if favoriteCount, err := models.NewFavoriteDaoInstance().QueryFavoriteCount(videoId); err == nil {
			c.JSON(http.StatusOK, Response{StatusCode: 0})
			video.FavoriteCount = favoriteCount + 1 //数据库数据加不上去+。+ ->？
			fmt.Println("favoriteCount add to : ", video.FavoriteCount)
		}
	} else if actionType == "2" {
		if favoriteCount, err := models.NewFavoriteDaoInstance().QueryFavoriteCount(videoId); err == nil {
			c.JSON(http.StatusOK, Response{StatusCode: 0})
			video.FavoriteCount = favoriteCount - 1
			fmt.Println("favoriteCount decrease to : ", video.FavoriteCount)
		} else {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "favorite_count decrease fail..."})
		}
	}
}

func FavoriteList(c *gin.Context) {
	//token := c.Query("token")
	userIdString := c.Query("user_id")
	userId, _ := strconv.Atoi(userIdString)
	videoIds := models.NewFavoriteDaoInstance().QueryFavoriteVideo(userId)
	videos := models.NewVideoDaoInstance().MQueryVideoByIds(videoIds)
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: videos,
	})
}
