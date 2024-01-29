package service

import (
	"begingo/common"
	"begingo/common/code"
	"begingo/common/log"
	"begingo/common/myerrors"
	"begingo/conf"
	"begingo/dao"
	"begingo/entity"
	"begingo/model"
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserSrv interface {
	Register(c *gin.Context, m *model.RegisterRequest) (int64, error)
	Login(c *gin.Context, m *model.LoginRequest) (*model.UserVO, error)
	Create(c *gin.Context, user *entity.User) (int64, error)
	Delete(c *gin.Context, req *common.DeleteRequest) (int64, error)
	Update(c *gin.Context, req *model.UserUpdateRequest) (int64, error)
	Get(c *gin.Context, userId int64) (*model.UserVO, error)
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

func (u *userService) Login(c *gin.Context, m *model.LoginRequest) (*model.UserVO, error) {
	byUser, err := u.dao.Users().Get(c, map[string]interface{}{"email": m.Email})
	if err != nil {
		log.Log().Error("查询失败: ", err)
		return nil, err
	}
	if byUser == nil {
		log.Log().Error("用户不存在")
		return nil, errors.New("用户不存在")
	}
	err = bcrypt.CompareHashAndPassword([]byte(byUser.Password), []byte(m.Password))
	if err != nil {
		log.Log().Error("密码错误: ", err)
		return nil, errors.New("密码错误")
	}
	userVO := &model.UserVO{
		ID:       byUser.ID,
		Nickname: byUser.Nickname,
		Email:    byUser.Email,
		Avatar:   byUser.Avatar,
		Gender:   byUser.Gender,
		UserRole: byUser.UserRole,
	}
	// 初始化session对象
	session := sessions.Default(c)
	// 设置session数据
	session.Set("userLogin", userVO.ID)
	err = session.Save()
	if err != nil {
		return nil, err
	}
	return userVO, err
}

func (u *userService) Create(c *gin.Context, user *entity.User) (int64, error) {
	// 判断邮箱是否注册
	byUser, err := u.dao.Users().Get(c, map[string]interface{}{"email": user.Email})
	if err != nil {
		log.Log().Info("验证邮箱: ", err)
	}
	if byUser != nil {
		log.Log().Info("邮箱已注册")
		return 0, errors.New("邮箱已注册")
	}
	userId, err := u.dao.Users().Create(c, user)
	return userId, err
}

func (u *userService) Delete(c *gin.Context, req *common.DeleteRequest) (int64, error) {
	// 根据 id 查询记录是否存在
	byUser, err := u.dao.Users().Get(c, map[string]interface{}{"id": req.Id})
	if err != nil {
		log.Log().Info("查询失败: ", err)
	}
	if byUser == nil {
		log.Log().Info("用户不存在")
		return 0, errors.New("用户不存在")
	}
	affectedRow, err := u.dao.Users().Delete(c, map[string]interface{}{"id": req.Id})
	return affectedRow, err
}

func (u *userService) Update(c *gin.Context, req *model.UserUpdateRequest) (int64, error) {
	if req == nil {
		return 0, myerrors.New("参数不能为空")
	}

	err := conf.Validate.Struct(req)
	if err != nil {
		return 0, err
	}

	// 根据 id 查询记录是否存在
	byUser, err := u.dao.Users().Get(c, map[string]interface{}{"id": req.ID})
	if err != nil {
		log.Log().Info("查询失败: ", err)
	}
	if byUser == nil {
		log.Log().Info("用户不存在")
		return 0, errors.New("用户不存在")
	}

	// 修改用户信息
	update := make(map[string]interface{})
	if len(req.Nickname) > 0 {
		update["nickname"] = req.Nickname
	}
	if len(req.Avatar) > 0 {
		update["avatar"] = req.Avatar
	}
	if req.Gender >= 0 {
		update["gender"] = req.Gender
	}
	if len(req.UserRole) > 0 {
		update["user_role"] = req.UserRole
	}
	affectedRow, err := u.dao.Users().UpdateCondition(c, map[string]interface{}{"id": req.ID}, update)
	if err != nil {
		return 0, err
	}
	return affectedRow, err
}

func (u *userService) Get(c *gin.Context, userId int64) (*model.UserVO, error) {

	user, err := u.dao.Users().Get(c, map[string]interface{}{"id": userId})
	if err != nil {
		return nil, err
	}
	userVO := &model.UserVO{
		ID:       user.ID,
		Nickname: user.Nickname,
		Email:    user.Email,
		Avatar:   user.Avatar,
		Gender:   user.Gender,
		UserRole: user.UserRole,
	}
	return userVO, err
}
