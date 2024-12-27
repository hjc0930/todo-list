package model

import (
	"time"
	"todo-list/consts"

	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	Id        int64      `gorm:"column:id;primary_key"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
	UserName  string     `gorm:"column:user_name;unique"`
	Password  string     `gorm:"column:password"`
}

func (*UserModel) TableName() string {
	return "user"
}

func (user *UserModel) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), consts.PasswordCost)

	if err != nil {
		return err
	}
	user.Password = string(bytes)

	return nil
}

func (user *UserModel) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
