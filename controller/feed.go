package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hjk-cloud/douyin/models"
	"github.com/hjk-cloud/douyin/service"
	"net/http"
	"strconv"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []*models.Video `json:"video_list,omitempty"`
	NextTime  int64           `json:"next_time,omitempty"`
}

func Feed(c *gin.Context) {
	token, flag := c.GetQuery("token")
	timeStamp := c.Query("latest_time")

	var latestTime time.Time
	times, err := strconv.ParseInt(timeStamp, 10, 64)
	if err == nil {
		latestTime = time.Unix(0, times*1e6)
	}
	var videos []*models.Video
	if !flag {
		token = "-"
	}

	videos, err = service.VideoListWithToken(token, latestTime)

	//fmt.Println("controller ----------", videos)
	if err == nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response:  Response{StatusCode: 0},
			VideoList: videos,
			NextTime:  time.Now().Unix(),
		})
	} else {
		c.JSON(http.StatusOK, FeedResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	}
}
