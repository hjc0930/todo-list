package com.hjc.todolist.service.impl;

import com.hjc.todolist.common.TaskStatusEnum;
import com.hjc.todolist.dao.TaskMapper;
import com.hjc.todolist.dto.CreateTaskDto;
import com.hjc.todolist.entity.Task;
import com.hjc.todolist.service.TaskService;
import com.hjc.todolist.util.PageQueryUtil;
import com.hjc.todolist.util.PageResult;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;

@Service
public class TaskServiceImpl implements TaskService {

    @Autowired
    private TaskMapper taskMapper;

    @Override
    public PageResult<List<Task>> getTaskList(PageQueryUtil pageUtil) {
        List<Task> taskList = taskMapper.findTaskList(pageUtil);
        int total = taskMapper.getTotalTasks(pageUtil);
        return new PageResult(pageUtil.getPage(), pageUtil.getPageSize(), total, taskList);
    }

    @Override
    public int addTask(CreateTaskDto createTaskDto) {
        Task task = new Task();
        task.setTaskName(createTaskDto.getTaskName());
        Optional.of(createTaskDto).map(CreateTaskDto::getTaskName).ifPresent(task::setTaskName);
        Optional.of(createTaskDto).map(CreateTaskDto::getTaskDesc).ifPresent(task::setTaskDescription);
        task.setUserId(1L);
        task.setStatus(TaskStatusEnum.COMPLETED.getCode());

        return taskMapper.insertSelective(task);
    }

    @Override
    public boolean deleteTask(Long taskId) {
        return false;
    }

    @Override
    public boolean updateTaskStatus(Long taskId, String status) {
        return false;
    }

    @Override
    public boolean updateTask(Task task) {
        return false;
    }
}
