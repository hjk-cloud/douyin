package controller

import (
	"github.com/RaymondCode/simple-demo/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type FeedResponse struct {
	models.Response
	VideoList []models.Video `json:"video_list,omitempty"`
	NextTime  int64          `json:"next_time,omitempty"`
}

func Feed(c *gin.Context) {
	videos := models.NewVideoDaoInstance().MQueryVideo()
	c.JSON(http.StatusOK, FeedResponse{
		Response:  models.Response{StatusCode: 0},
		VideoList: videos,
		NextTime:  time.Now().Unix(),
	})
}
