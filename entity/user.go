package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID       int64          `json:"id" gorm:"column:id;primary_key"`
	Nickname string         `json:"nickname" gorm:"column:nickname"`
	Email    string         `json:"email" gorm:"column:email"`
	Password string         `json:"password,omitempty" gorm:"column:password"`
	Avatar   string         `json:"avatar,omitempty" gorm:"column:avatar"`
	Gender   int            `json:"gender" gorm:"column:gender"`
	UserRole string         `json:"user_role,omitempty" gorm:"column:user_role"`
	CreateAt time.Time      `json:"create_at" gorm:"column:create_at"`
	UpdateAt time.Time      `json:"update_at" gorm:"column:update_at"`
	DeleteAt gorm.DeletedAt `json:"delete_at" gorm:"column:delete_at"`
}
