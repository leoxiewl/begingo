package dao

import (
	"begingo/entity"
	"context"
)

type UserDao interface {
	Create(ctx context.Context, user *entity.User) (int64, error)
	UpdateCondition(ctx context.Context, where map[string]interface{}, update map[string]interface{}) (int64, error)
	Get(ctx context.Context, where map[string]interface{}) (*entity.User, error)
}
