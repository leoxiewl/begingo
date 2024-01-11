package dao

import (
	"begingo/entity"
	"context"
)

type UserDao interface {
	Create(ctx context.Context, user *entity.User) error
}
