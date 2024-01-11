package service

import (
	"begingo/dao"
	"begingo/entity"
	"begingo/model"
	"context"
	"time"
)

type UserSrv interface {
	Create(ctx context.Context, req *model.UserAddRequest) error
}
type userService struct {
	dao dao.Factory
}

func newUsers(srv *service) *userService {
	return &userService{dao: srv.store}
}

func (u *userService) Create(ctx context.Context, req *model.UserAddRequest) error {
	// 参数校验

	user := &entity.User{
		Nickname: req.Nickname,
		Email:    req.Email,
		Password: req.Password,
		Avatar:   req.Avatar,
		Gender:   req.Gender,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}
	if err := u.dao.Users().Create(ctx, user); err != nil {
		return err
	}
	return nil
}
