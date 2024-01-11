package model

type UserAddRequest struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
	Gender   int    `json:"gender"`
}
