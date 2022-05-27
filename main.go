package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hjk-cloud/douyin/models"
	"github.com/hjk-cloud/douyin/routers"
)

func main() {
	models.Init()

	r := gin.Default()

	routers.InitRouter(r)

	r.Run()
}
