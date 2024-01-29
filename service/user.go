package service

import (
	"begingo/common/code"
	"begingo/common/log"
	"begingo/dao"
	"begingo/entity"
	"begingo/model"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserSrv interface {
	Register(c *gin.Context, m *model.RegisterRequest) (int64, error)
}
type userService struct {
	dao dao.Factory
}

func newUsers(srv *service) *userService {
	return &userService{dao: srv.store}
}

func (u *userService) Register(c *gin.Context, m *model.RegisterRequest) (int64, error) {
	// 参数校验

	// 判断邮箱是否注册
	byUser, err := u.dao.Users().Get(c, map[string]interface{}{"email": m.Email})
	if err != nil {
		log.Log().Info("验证邮箱: ", err)
	}
	if byUser != nil {
		log.Log().Info("邮箱已注册")
		return 0, errors.New("邮箱已注册")
	}
	// 生成密码
	bytes, err := bcrypt.GenerateFromPassword([]byte(m.Password), code.PassWordCost)
	if err != nil {
		return 0, err
	}
	var user *entity.User
	user = &entity.User{
		Email:    m.Email,
		Password: string(bytes),
		Nickname: m.Nickname,
	}
	// 保存用户信息
	userId, err := u.dao.Users().Create(c, user)
	return userId, err
}
