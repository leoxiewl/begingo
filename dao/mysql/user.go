package mysql

import (
	"begingo/entity"
	"context"
	"gorm.io/gorm"
)

type users struct {
	db *gorm.DB
}

func newUsers(ds *datastore) *users {
	return &users{db: ds.db}
}

func (u *users) Create(ctx context.Context, user *entity.User) error {
	return u.db.Create(&user).Error
}
