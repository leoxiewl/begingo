package v1

import (
	"begingo/common"
	"begingo/common/code"
	"begingo/common/response"
	"begingo/conf"
	"begingo/dao"
	"begingo/entity"
	"begingo/model"
	srv "begingo/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"strconv"
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

func (u *UserHandler) Login(c *gin.Context) {
	var req model.LoginRequest

	// 绑定参数
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failed(c, code.ErrBind, "参数绑定失败")
		return
	}

	// 校验参数
	err := conf.Validate.Struct(req)
	if err != nil {
		response.Failed(c, code.ErrValidation, err.Error())
		return
	}

	user, err := u.srv.Users().Login(c, &req)
	if err != nil {
		response.Failed(c, code.ErrCommon, err.Error())
		return
	}
	response.Success(c, code.SucCommon, user)
}

func (u *UserHandler) GetUserLogin(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("currentUser")
	if user == nil {
		response.Failed(c, code.ErrCommon, "用户未登录")
	}
	userId := user.(model.UserVO).ID
	if userId <= 0 {
		response.Failed(c, code.ErrCommon, "用户未登录")
	}
	// 查库获取用户信息，保证信息是最新的
	user, err := u.srv.Users().Get(c, userId)
	if err != nil {
		response.Failed(c, code.ErrCommon, err.Error())
		return
	}
	response.Success(c, code.SucCommon, user)
}

func (u *UserHandler) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("currentUser")
	err := session.Save()
	if err != nil {
		response.Failed(c, code.ErrCommon, err.Error())
		return
	}
	response.Success(c, code.SucCommon, nil)
}

// Create 当作代码示例，密码没有做加密处理
func (u *UserHandler) Create(c *gin.Context) {
	var req entity.User

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failed(c, code.ErrBind, err.Error())
		return
	}

	err := conf.Validate.Struct(req)
	if err != nil {
		response.Failed(c, code.ErrValidation, err.Error())
		return
	}

	userId, err := u.srv.Users().Create(c, &req)
	if err != nil {
		response.Failed(c, code.ErrCommon, err.Error())
		return
	}
	response.Success(c, code.SucCommon, userId)

}

func (u *UserHandler) Delete(c *gin.Context) {
	var req common.DeleteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failed(c, code.ErrBind, err.Error())
		return
	}

	affectedRow, err := u.srv.Users().Delete(c, &req)
	if err != nil {
		response.Failed(c, code.ErrCommon, err.Error())
		return
	}
	response.Success(c, code.SucCommon, affectedRow)
}

func (u *UserHandler) Update(c *gin.Context) {
	var req model.UserUpdateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failed(c, code.ErrBind, err.Error())
		return
	}

	// 参数校验
	err := conf.Validate.Struct(req)
	if err != nil {
		response.Failed(c, code.ErrValidation, err.Error())
		return
	}

	affectedRow, err := u.srv.Users().Update(c, &req)
	if err != nil {
		response.Failed(c, code.ErrCommon, err.Error())
		return
	}
	response.Success(c, code.SucCommon, affectedRow)
}

func (u *UserHandler) Get(c *gin.Context) {

	idParam := c.Query("id")

	// Convert id to int64
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return
	}
	user, err := u.srv.Users().Get(c, id)
	if err != nil {
		response.Failed(c, code.ErrCommon, err.Error())
		return
	}
	response.Success(c, code.SucCommon, user)

}

func (u *UserHandler) ListPage(c *gin.Context) {
	var req model.UserQueryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failed(c, code.ErrBind, "参数绑定失败: "+err.Error())
		return
	}

	pageRes, err := u.srv.Users().ListPage(c, &req)
	if err != nil {
		response.Failed(c, code.ErrCommon, err.Error())
		return
	}
	response.Success(c, code.SucCommon, pageRes)
}
