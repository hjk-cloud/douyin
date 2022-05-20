package test

import (
	"github.com/RaymondCode/simple-demo/models"
	"github.com/gin-gonic/gin"
)

func UserList(c *gin.Context) {
	models.GetUserList()
}
