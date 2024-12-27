package service

import (
	"context"
	"sync"
	"time"
	"todo-list/pkg/ctl"
	"todo-list/pkg/utils"
	"todo-list/repository/db/dao"
	"todo-list/repository/db/model"
	"todo-list/types"
)

type TaskSrv struct{}

var TaskOnce sync.Once
var TaskSrvInstance *TaskSrv

func GetTaskSrv() *TaskSrv {
	TaskOnce.Do(func() {
		TaskSrvInstance = &TaskSrv{}
	})

	return TaskSrvInstance
}

func (s *TaskSrv) CreateTask(ctx context.Context, req *types.CreateTaskReq) (resp interface{}, err error) {
	userInfo, err := ctl.GetUserInfo(ctx)
	if err != nil {
		utils.LogrusObj.Error(err)
		return
	}

	currentUser, err := dao.NewUserDao(ctx).FindUserByUserId(userInfo.Id)
	if err != nil {
		utils.LogrusObj.Info(err)
		return
	}
	task := &model.TaskModel{
		Uid:       currentUser.Id,
		Title:     req.Title,
		Content:   req.Content,
		Status:    0,
		StartTime: time.Now().Unix(),
	}

	err = dao.NewTaskDao(ctx).CreateTask(task)
	if err != nil {
		utils.LogrusObj.Info(err)
		return
	}
	return ctl.RespSuccess(), nil
}
