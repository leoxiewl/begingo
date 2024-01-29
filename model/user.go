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
