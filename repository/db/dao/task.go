package dao

import (
	"context"

	"gorm.io/gorm"

	"todo-list/repository/db/model"
	"todo-list/types"
)

type TaskDao struct {
	*gorm.DB
}

func NewTaskDao(ctx context.Context) *TaskDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &TaskDao{NewDBClient(ctx)}
}

// CreateTask 创建Task
func (s *TaskDao) CreateTask(task *model.TaskModel) error {
	return s.Model(&model.TaskModel{}).Create(&task).Error
}

// ListTask List Task的情况
func (s *TaskDao) ListTask(start, limit int, userId int64) (result []*model.TaskModel, total int64, err error) {
	err = s.DB.Model(&model.TaskModel{}).
		Where("uid = ?", userId).
		Count(&total).Error

	if err != nil {
		return
	}
	err = s.DB.Model(&model.TaskModel{}).
		Where("uid = ?", userId).
		Limit(limit).
		Offset((start - 1) * limit).
		Find(&result).Error

	return
}

// FindTaskByIdAndUserId 通过id和user_id找到task
func (s *TaskDao) FindTaskByIdAndUserId(id, uId int64) (r *model.TaskModel, err error) {
	err = s.DB.Model(&model.TaskModel{}).Where("id = ? AND uid = ?", id, uId).First(&r).Error
	return
}

// UpdateTask 修改
func (s *TaskDao) UpdateTask(uId int64, req *types.UpdateTaskReq) error {
	var task model.TaskModel
	err := s.DB.Model(&model.TaskModel{}).
		Where("id = ? AND uid=?", req.Id, uId).
		First(&task).Error

	if err != nil {
		return err
	}
	if req.Status != 0 {
		task.Status = req.Status
	}
	if req.Title != "" {
		task.Title = req.Title
	}
	if req.Content != "" {
		task.Content = req.Content
	}

	return s.Save(&task).Error
}

// SearchTask 搜索Task
func (s *TaskDao) SearchTask(uId int64, info string) (tasks []*model.TaskModel, err error) {
	err = s.DB.Model(&model.TaskModel{}).
		Where("uid = ? AND (title LIKE ? OR content LIKE ?)",
			uId,
			"%"+info+"%", "%"+info+"%").
		Find(&tasks).
		Error
	return
}

// DeleteTaskById 通过id删除
func (s *TaskDao) DeleteTaskById(tId, uId int64) (err error) {
	err = s.DB.Model(&model.TaskModel{}).
		Where("id = ? AND uid = ?", tId, uId).
		Delete(&model.TaskModel{}).Error
	return err
}
