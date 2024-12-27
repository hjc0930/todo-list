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

func (s *TaskSrv) ListTask(ctx context.Context, req *types.ListTasksReq) (resp interface{}, err error) {
	userInfo, err := ctl.GetUserInfo(ctx)
	if err != nil {
		utils.LogrusObj.Error(err)
		return
	}
	tasks, total, err := dao.NewTaskDao(ctx).ListTask(req.Start, req.Limit, userInfo.Id)
	tRespList := make([]*types.ListTasksResp, 0)
	for _, task := range tasks {
		tRespList = append(tRespList, &types.ListTasksResp{
			Id:      task.Id,
			Title:   task.Title,
			Content: task.Content,
			View:    task.View(),
			Status:  task.Status,
			Created: task.CreatedAt.Unix(),
			EndTime: task.EndTime,
		})
	}

	return ctl.RespList(tRespList, total), nil
}

func (s *TaskSrv) DetailsTask(ctx context.Context, req *types.DetailsTaskReq) (resp interface{}, err error) {
	userInfo, err := ctl.GetUserInfo(ctx)
	if err != nil {
		utils.LogrusObj.Error(err)
		return
	}
	task, err := dao.NewTaskDao(ctx).FindTaskByIdAndUserId(req.Id, userInfo.Id)
	if err != nil {
		utils.LogrusObj.Info(err)
		return
	}

	task.AddView()
	tResp := &types.ListTasksResp{
		Id:        task.Id,
		Title:     task.Title,
		Content:   task.Content,
		View:      task.View(),
		Status:    task.Status,
		Created:   task.CreatedAt.Unix(),
		StartTime: task.StartTime,
		EndTime:   task.EndTime,
	}
	return ctl.RespSuccessWithData(tResp), nil
}

func (s *TaskSrv) UpdateTask(ctx context.Context, req *types.UpdateTaskReq) (resp interface{}, err error) {
	userInfo, err := ctl.GetUserInfo(ctx)
	if err != nil {
		utils.LogrusObj.Error(err)
		return
	}
	err = dao.NewTaskDao(ctx).UpdateTask(userInfo.Id, req)

	if err != nil {
		utils.LogrusObj.Info(err)
		return
	}
	return ctl.RespSuccess(), nil
}

func (s *TaskSrv) SearchTask(ctx context.Context, req *types.SearchTaskReq) (resp interface{}, err error) {
	userInfo, err := ctl.GetUserInfo(ctx)
	if err != nil {
		utils.LogrusObj.Error(err)
		return
	}
	tasks, err := dao.NewTaskDao(ctx).SearchTask(userInfo.Id, req.Info)
	if err != nil {
		utils.LogrusObj.Info(err)
		return
	}

	tRespList := make([]*types.ListTasksResp, 0)
	for _, task := range tasks {
		tRespList = append(tRespList, &types.ListTasksResp{
			Id:      task.Id,
			Title:   task.Title,
			Content: task.Content,
			View:    task.View(),
			Status:  task.Status,
			Created: task.CreatedAt.Unix(),
			EndTime: task.EndTime,
		})
	}
	return ctl.RespSuccessWithData(tRespList), nil
}

func (s *TaskSrv) DeleteTask(ctx context.Context, req *types.DeleteTaskReq) (resp interface{}, err error) {
	userInfo, err := ctl.GetUserInfo(ctx)
	if err != nil {
		utils.LogrusObj.Error(err)
		return
	}
	err = dao.NewTaskDao(ctx).DeleteTaskById(req.Id, userInfo.Id)
	if err != nil {
		utils.LogrusObj.Error(err)
		return
	}
	return ctl.RespSuccess(), nil
}
