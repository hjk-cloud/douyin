package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hjk-cloud/douyin/define"
	"github.com/hjk-cloud/douyin/models"
	"github.com/hjk-cloud/douyin/util"
	"github.com/hjk-cloud/douyin/util/jwt"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"
)

type PublishFlow struct {
	Token      string
	Title      string
	Data       *multipart.FileHeader
	c          *gin.Context
	Video      *models.Video
	AuthorId   int
	PlayUrl    string
	CoverUrl   string
	SubmitTime time.Time
}

func Publish(token string, title string, data *multipart.FileHeader, c *gin.Context) error {
	return NewPublishFlow(token, title, data, c).Do()
}

func NewPublishFlow(token string, title string, data *multipart.FileHeader, c *gin.Context) *PublishFlow {
	return &PublishFlow{Token: token, Title: title, Data: data, c: c}
}

func (f *PublishFlow) Do() error {
	if err := f.checkParam(); err != nil {
		return err
	}
	if err := f.prepareData(); err != nil {
		return err
	}
	if err := f.packData(); err != nil {
		return err
	}
	return nil
}

func (f *PublishFlow) checkParam() error {

	return nil
}

func (f *PublishFlow) prepareData() error {
	userId, err := jwt.JWTAuth(f.Token)
	if err != nil {
		return err
	}
	//fmt.Println("prepareData----UserId", userId)
	f.AuthorId = userId
	filename := filepath.Base(f.Data.Filename)
	finalName := fmt.Sprintf("%d_%s", userId, filename)
	saveFile := filepath.Join("./public/", finalName)
	if err := f.c.SaveUploadedFile(f.Data, saveFile); err != nil {
		return err
	}
	if _, err := util.GetSnapshot(define.StaticSourcePath+finalName, define.StaticSourcePath+"pic\\"+strings.Split(finalName, ".")[0], 1); err != nil {
		return err
	}
	playUrl := define.URL + "/static/" + finalName
	picUrl := define.URL + "/static/pic/" + strings.Split(finalName, ".")[0] + ".png"
	f.PlayUrl = playUrl
	f.CoverUrl = picUrl
	f.SubmitTime = time.Now().Local()
	return nil
}

func (f *PublishFlow) packData() error {
	f.Video = &models.Video{
		AuthorId:   f.AuthorId,
		PlayUrl:    f.PlayUrl,
		CoverUrl:   f.CoverUrl,
		Title:      f.Title,
		SubmitTime: f.SubmitTime,
	}
	//fmt.Println("packData----time", f.SubmitTime)
	videoDao := models.NewVideoDaoInstance()
	if err := videoDao.PublishVideo(f.Video); err != nil {
		return err
	}
	return nil
}
