package model

import "time"

type UserModel struct {
	Id             int64      `gorm:"column:id;primary_key"`
	CreatedAt      *time.Time `gorm:"column:created_at"`
	UpdatedAt      *time.Time `gorm:"column:updated_at"`
	DeletedAt      *time.Time `gorm:"column:deleted_at"`
	UserName       string     `gorm:"column:user_name;unique"`
	PasswordDigest string     `gorm:"column:password_digest"`
}

func (*UserModel) TableName() string {
	return "user"
}
