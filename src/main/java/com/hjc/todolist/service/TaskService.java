package com.hjc.todolist.service;

import com.hjc.todolist.dto.CreateTaskDto;
import com.hjc.todolist.entity.Task;
import com.hjc.todolist.util.PageQueryUtil;
import com.hjc.todolist.util.PageResult;

public interface TaskService {

    /**
     * 获取任务列表
     *
     * @param pageUtil 分页工具类
     * @return 任务列表
     */
    PageResult getTaskList(PageQueryUtil pageUtil);
    /**
     * 添加任务
     *
     * @param createTaskDto 任务名称
     * @return 任务ID
     */
    int addTask(CreateTaskDto createTaskDto);

    /**
     * 删除任务
     *
     * @param taskId 任务ID
     * @return 是否成功
     */
    boolean deleteTask(Long taskId);

    /**
     * 更新任务状态
     *
     * @param taskId  任务ID
     * @param status  新状态
     * @return 是否成功
     */
    boolean updateTaskStatus(Long taskId, String status);

    /**
     * 更新任务
     *
     * @param task 任务ID
     * @return 是否成功
     */
    boolean updateTask(Task task);
}
