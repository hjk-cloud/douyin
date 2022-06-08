package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hjk-cloud/douyin/models"
	"github.com/hjk-cloud/douyin/routers"
	"github.com/hjk-cloud/douyin/util"
)

func main() {
	models.Init()

	r := gin.Default()

	routers.InitRouter(r)

	//用来测试redis连接
	util.GetAll()

	r.Run()
}
