package com.hjc.todolist.service.impl;

import com.hjc.todolist.common.TaskStatusEnum;
import com.hjc.todolist.dao.TaskMapper;
import com.hjc.todolist.dao.UserMapper;
import com.hjc.todolist.dto.CreateTaskDto;
import com.hjc.todolist.entity.Task;
import com.hjc.todolist.entity.User;
import com.hjc.todolist.service.TaskService;
import com.hjc.todolist.util.BeanUtil;
import com.hjc.todolist.util.PageQueryUtil;
import com.hjc.todolist.util.PageResult;
import com.hjc.todolist.vo.TaskListVo;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.List;

@Service
public class TaskServiceImpl implements TaskService {

    @Autowired
    private TaskMapper taskMapper;
    @Autowired
    private UserMapper userMapper;

    @Override
    public PageResult<TaskListVo> getTaskList(PageQueryUtil pageUtil) {
        List<Task> taskList = taskMapper.findTaskList(pageUtil);
        List<Long> taskUserIds = taskList.stream().map(Task::getUserId).toList();
        List<User> userList = userMapper.findUserByIds(taskUserIds);
        int total = taskMapper.getTotalTasks(pageUtil);

        ArrayList<TaskListVo> taskVoList = new ArrayList<>();
        taskList.forEach(task -> {
            TaskListVo taskListVo = new TaskListVo();
            BeanUtil.copyProperties(task, taskListVo);

            userList.stream()
                    .filter(u -> u.getUserId().equals(task.getUserId()))
                    .findFirst()
                    .map(User::getUserName)
                    .ifPresent(taskListVo::setUserName);
            taskVoList.add(taskListVo);
        });

        PageResult<TaskListVo> pageResult = new PageResult<>();
        pageResult.setPage(pageUtil.getPage());
        pageResult.setPageSize(pageUtil.getPageSize());
        pageResult.setTotalCount(total);
        pageResult.setList(taskVoList);

        return pageResult;
    }

    @Override
    public int addTask(CreateTaskDto createTaskDto) {
        Task task = new Task();
        BeanUtil.copyProperties(createTaskDto, task);
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
