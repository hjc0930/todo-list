package service

import (
	"context"
	"sync"
	"todo-list/repository/db/dao"
	"todo-list/repository/db/model"
	"todo-list/types"

	"gorm.io/gorm"
)

var UserSrvIns *UserSrv
var UserSrvOnce sync.Once

type UserSrv struct{}

func GetUserSrv() *UserSrv {
	UserSrvOnce.Do(func() {
		UserSrvIns = &UserSrv{}
	})
	return UserSrvIns
}

func (s *UserSrv) UserRegister(ctx context.Context, req *types.UserRegisterReq) (resp interface{}, err error) {
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.FindUserByUserName(req.UserName)

	switch err {
	case gorm.ErrRecordNotFound:
		user = &model.UserModel{}
	}
	// userDao := dao.NewUserDao(ctx);
	// return userDao
}
