package model

import (
	"begingo/common"
	"time"
)

type UserAddRequest struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
	Gender   int    `json:"gender"`
}

type RegisterRequest struct {
	Nickname      string `json:"nickname" validate:"required,min=2,max=20"`
	Email         string `json:"email" validate:"required,email"`
	Password      string `json:"password" validate:"required,min=6,max=40"`
	CheckPassword string `json:"check_password" validate:"required,eqfield=Password"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=40"`
}

type UserUpdateRequest struct {
	ID       int64  `json:"id" validate:"required"`
	Nickname string `json:"nickname" validate:"omitempty,min=2,max=20"`
	Avatar   string `json:"avatar" validate:"omitempty,url"`
	Gender   int    `json:"gender" validate:"omitempty,oneof=0 1 2"`
	UserRole string `json:"user_role" validate:"omitempty"`
}

type UserQueryRequest struct {
	ID          int64              `json:"id"`
	Nickname    string             `json:"nickname" validate:"omitempty,min=0,max=20"`
	Email       string             `json:"email" validate:"omitempty,email"`
	UserRole    string             `json:"user_role" validate:"omitempty,oneof=admin user"`
	CreateAt    string             `json:"create_at" `
	PageRequest common.PageRequest `json:"page_request"`
}

type UserVO struct {
	ID       int64     `json:"id"`
	Nickname string    `json:"nickname"`
	Email    string    `json:"email" `
	Avatar   string    `json:"avatar" `
	Gender   int       `json:"gender" `
	UserRole string    `json:"user_role"`
	CreateAt time.Time `json:"create_at" `
}
