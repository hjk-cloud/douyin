package main

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/router"
	"github.com/gin-gonic/gin"
)

func main() {
	model.Init()

	r := gin.Default()

	router.InitRouter(r)

	r.Run()
}
