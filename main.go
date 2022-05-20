package main

import (
	"github.com/RaymondCode/simple-demo/models"
	"github.com/RaymondCode/simple-demo/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	models.Init()

	r := gin.Default()

	routers.InitRouter(r)

	r.Run()
}
