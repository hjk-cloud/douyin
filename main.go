package main

import (
	"github.com/RaymondCode/simple-demo/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routers.InitRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
