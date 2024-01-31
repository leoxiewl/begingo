package mysql

import (
	"begingo/common/log"
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

func (u *users) Create(ctx context.Context, user *entity.User) (int64, error) {
	result := u.db.Omit("create_at", "update_at").Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}
	return user.ID, nil
}

func (u *users) Delete(ctx context.Context, where map[string]interface{}) (int64, error) {
	result := u.db.Where(where).Delete(&entity.User{})
	if result.Error != nil {
		log.Log().Error("delete user myerrors: ", result.Error)
	}
	return result.RowsAffected, result.Error
}

func (u *users) UpdateCondition(ctx context.Context, where map[string]interface{}, update map[string]interface{}) (int64, error) {
	result := u.db.Model(&entity.User{}).Where(where).Updates(update)
	if result.Error != nil {
		log.Log().Error("update user myerrors: ", result.Error)
	}
	return result.RowsAffected, result.Error
}

func (u *users) Get(ctx context.Context, where map[string]interface{}) (*entity.User, error) {
	var user entity.User
	err := u.db.Where(where).First(&user).Error
	if err != nil {
		log.Log().Error("get user myerrors: ", err)
		return nil, err
	}
	return &user, nil
}

func (u *users) ListPage(ctx context.Context, where map[string]interface{}, page int, pageSize int) ([]*entity.User, error) {
	var users []*entity.User
	query := u.db.Model(&entity.User{})
	for key, value := range where {
		query = query.Where(key, value)
	}
	err := query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&users).Error
	if err != nil {
		log.Log().Error("list user errors: ", err)
		return nil, err
	}
	return users, nil
}

func (u *users) Count(ctx context.Context, where map[string]interface{}) (int64, error) {
	var count int64
	query := u.db.Model(&entity.User{})
	for key, value := range where {
		query = query.Where(key, value)
	}
	err := query.Count(&count).Error
	if err != nil {
		log.Log().Error("count user errors: ", err)
		return 0, err
	}
	return count, nil
}
