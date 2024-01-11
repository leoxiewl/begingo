package api

import (
	"begingo/dao"
	"begingo/model"
	"begingo/pkg/code"
	"begingo/util"
	"github.com/gin-gonic/gin"

	srv "begingo/service"
)

type UserHandler struct {
	srv   srv.Service
	store dao.Factory
}

// NewUserHandler creates a user handler.
func NewUserHandler(store dao.Factory) *UserHandler {
	return &UserHandler{
		srv:   srv.NewService(store),
		store: store,
	}
}

func (u *UserHandler) Create(c *gin.Context) {
	var req model.UserAddRequest

	// 绑定参数
	if err := c.ShouldBindJSON(&req); err != nil {
		util.Failed(c, code.ErrBind, "参数绑定失败")
		return
	}

	// 校验参数 TODO

	err := u.srv.Users().Create(c, &req)
	if err != nil {
		util.Failed(c, code.ErrSuccess, "创建用户失败")
		return
	}
	util.Success(c, 200, "用户信息")
}
