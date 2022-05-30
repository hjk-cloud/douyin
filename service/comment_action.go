package service

import (
	"fmt"
	"github.com/hjk-cloud/douyin/models"
	"github.com/hjk-cloud/douyin/util/jwt"
	"time"
)

type CommentActionFlow struct {
	UserId      int
	Token       string
	VideoId     int
	ActionType  string
	CommentText string
	commentId   int
	Comment     *models.Comment
}

func CommentAction(userId int, token string, videoId int, actionType string, commentText string, commentId int) (*models.Comment, error) {
	return NewCommentActionFlow(userId, token, videoId, actionType, commentText, commentId).Do()
}

func NewCommentActionFlow(userId int, token string, videoId int, actionType string, commentText string, commentId int) *CommentActionFlow {
	return &CommentActionFlow{
		UserId:      userId,
		Token:       token,
		VideoId:     videoId,
		ActionType:  actionType,
		CommentText: commentText,
		commentId:   commentId,
	}
}

func (f *CommentActionFlow) Do() (*models.Comment, error) {
	if err := f.checkParam(); err != nil {
		return nil, err
	}
	if err := f.prepareData(); err != nil {
		return nil, err
	}
	if err := f.packData(); err != nil {
		return nil, err
	}
	return f.Comment, nil

}

func (f *CommentActionFlow) checkParam() error {

	return nil
}

//遇到了老问题，前端传来的user_id为0，只能通过token拿user_id
func (f *CommentActionFlow) prepareData() error {
	userId, err := jwt.JWTAuth(f.Token)
	if err != nil {
		return err
	}
	f.UserId = userId
	return nil
}

func (f *CommentActionFlow) packData() error {
	comment := &models.Comment{
		UserId:  f.UserId,
		VideoId: f.VideoId,
	}
	commentDao := models.NewCommentDaoInstance()
	userDao := models.NewUserDaoInstance()
	relationDao := models.NewRelationDaoInstance()

	if f.ActionType == "1" {
		comment.Content = f.CommentText
		comment.CreateDate = time.Now().Format("01-02")
		//fmt.Println("service-----", comment.CreateDate)
		if err := commentDao.CreateComment(comment); err != nil {
			return err
		}
		user, err := userDao.QueryUserById(comment.UserId)
		if err != nil {
			return err
		}
		user.IsFollow = relationDao.QueryRelationState(f.UserId, user.Id)
		user.FollowerCount, _ = relationDao.QueryRelationCountByToUserId(user.Id)
		user.FollowCount, _ = relationDao.QueryRelationCountByUserId(user.Id)
		comment.User = *user
	} else {
		comment.Id = f.commentId
		if err := commentDao.DeleteComment(comment); err != nil {
			return err
		}
	}
	f.Comment = comment
	fmt.Println(f.Comment)
	return nil
}
