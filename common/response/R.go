package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 统一返回类在 Controller 包下，不需要引入其他包

type R struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(c *gin.Context, code int, data interface{}) {
	c.JSON(200, R{
		Code: code,
		Msg:  "success",
		Data: data,
	})
}

func Failed(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, R{
		Code: code,
		Msg:  msg,
		Data: "",
	})
}
