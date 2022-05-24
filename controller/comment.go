package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CommentListResponse struct {
	models.Response
	CommentList []models.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	models.Response
	Comment models.Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	actionType := c.Query("action_type")

	if user, exist := usersLoginInfo[token]; exist {
		if actionType == "1" {
			text := c.Query("comment_text")
			c.JSON(http.StatusOK, CommentActionResponse{Response: models.Response{StatusCode: 0},
				Comment: models.Comment{
					Id:         1,
					User:       user,
					Content:    text,
					CreateDate: "05-01",
				}})
			return
		}
		c.JSON(http.StatusOK, models.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, models.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

func CommentList(c *gin.Context) {
	//token := c.Query("token")
	videoIdString := c.Query("video_id")
	videoId, _ := strconv.Atoi(videoIdString)
	fmt.Println("videoId ----------", videoId)
	comments := models.NewCommentDaoInstance().MQueryCommentById(videoId)
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    models.Response{StatusCode: 0},
		CommentList: comments,
	})
}
