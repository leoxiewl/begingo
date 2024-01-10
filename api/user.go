package api

import (
	"begingo/util"
	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	util.Success(c, 200, "用户信息")
}

func GetUserList(c *gin.Context) {
	util.Failed(c, 400, "没有用户列表")
}
