package model

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

type UserVO struct {
	ID       int64  `json:"id"`
	Nickname string `json:"nickname"`
	Email    string `json:"email" `
	Avatar   string `json:"avatar" `
	Gender   int    `json:"gender" `
	UserRole string `json:"user_role"`
}

type UserUpdateRequest struct {
	ID       int64  `json:"id" validate:"required"`
	Nickname string `json:"nickname" validate:"omitempty,min=2,max=20"`
	Avatar   string `json:"avatar" validate:"omitempty,url"`
	Gender   int    `json:"gender" validate:"omitempty,oneof=0 1 2"`
	UserRole string `json:"user_role" validate:"omitempty"`
}
