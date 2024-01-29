package v1

import (
	"begingo/common/code"
	"begingo/common/response"
	"begingo/conf"
	"begingo/dao"
	"begingo/model"
	srv "begingo/service"
	"fmt"
	"github.com/gin-gonic/gin"
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

func (u *UserHandler) Register(c *gin.Context) {
	var req model.RegisterRequest

	// 绑定参数
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failed(c, code.ErrBind, "参数绑定失败")
		return
	}

	// 校验参数
	err := conf.Validate.Struct(req)
	if err != nil {
		fmt.Println(err.Error())

		//if _, ok := err.(*validator.InvalidValidationError); ok {
		//	fmt.Println(err)
		//}
		response.Failed(c, code.ErrValidation, err.Error())
		return
	}

	userId, err := u.srv.Users().Register(c, &req)
	if err != nil {
		response.Failed(c, code.ErrCommon, err.Error())
		return
	}
	response.Success(c, code.SucCommon, userId)
}