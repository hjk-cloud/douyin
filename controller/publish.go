package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

type VideoListResponse struct {
	models.Response
	VideoList []models.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")

	user, err := models.NewUserDaoInstance().QueryUserByToken(token)
	if err != nil {
		c.JSON(http.StatusOK, models.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	videos := models.NewVideoDaoInstance().MQueryVideo()
	c.JSON(http.StatusOK, VideoListResponse{
		Response: models.Response{
			StatusCode: 0,
		},
		VideoList: videos,
	})
}
