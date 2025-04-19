package com.hjc.todolist.service;

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
     * @param taskName 任务名称
     * @param taskDesc 任务描述
     * @return 任务ID
     */
    Long addTask(String taskName, String taskDesc);

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
