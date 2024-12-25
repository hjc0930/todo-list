package service

import (
	"context"
	"sync"
	"todo-list/types"
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

func (s *UserSrv) UserRegister(c context.Context, req *types.UserRegisterReq) (resp interface{}, err error) {

	return
}
