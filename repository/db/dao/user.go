package dao

import (
	"context"

	"gorm.io/gorm"

	"todo-list/repository/db/model"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &UserDao{NewDBClient(ctx)}
}

// FindUserByUserName 根据用户名找到用户
func (dao *UserDao) FindUserByUserName(userName string) (user *model.UserModel, err error) {
	err = dao.DB.Model(&model.UserModel{}).Where("user_name=?", userName).
		First(&user).Error

	return
}

// FindUserByUserId 根据用户id找到用户
func (dao *UserDao) FindUserByUserId(id int64) (user *model.UserModel, err error) {
	err = dao.DB.Model(&model.UserModel{}).Where("id=?", id).
		First(&user).Error

	return
}

// CreateUser 创建User
func (dao *UserDao) CreateUser(user *model.UserModel) (err error) {
	err = dao.DB.Model(&model.UserModel{}).Create(user).Error
	return
}

func (s *UserDao) Update(in *model.UserModel) error {
	return s.DB.Model(&model.UserModel{}).Updates(in).Error
}
