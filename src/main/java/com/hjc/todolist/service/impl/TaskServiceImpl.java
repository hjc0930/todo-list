package com.hjc.todolist.service.impl;

import com.hjc.todolist.dao.TaskMapper;
import com.hjc.todolist.entity.Task;
import com.hjc.todolist.service.TaskService;
import com.hjc.todolist.util.PageQueryUtil;
import com.hjc.todolist.util.PageResult;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class TaskServiceImpl implements TaskService {

    @Autowired
    private TaskMapper taskMapper;

    @Override
    public PageResult getTaskList(PageQueryUtil pageUtil) {
        List<Task> taskList = taskMapper.findTaskList(pageUtil);
        int total = taskMapper.getTotalTasks(pageUtil);
        new PageResult(pageUtil.getPage(), pageUtil.getPageSize(), total, taskList);
        return null;
    }

    @Override
    public Long addTask(String taskName, String taskDesc) {
        return 0L;
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
