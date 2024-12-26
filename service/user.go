package service

import (
	"context"
	"errors"
	"sync"
	"todo-list/pkg/ctl"
	"todo-list/pkg/utils"
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
		user = &model.UserModel{
			UserName: req.UserName,
		}
		// Password encrypted storage
		if err = user.SetPassword(req.Password); err != nil {
			utils.LogrusObj.Error(err)
			return
		}

		if err = userDao.CreateUser(user); err != nil {
			utils.LogrusObj.Error(err)
			return
		}

		return ctl.RespSuccess(), nil
	case nil:
		err = errors.New("user already exists")
		return
	default:
		return
	}
}
