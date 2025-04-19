package com.hjc.todolist.dao;

import com.hjc.todolist.entity.Task;
import com.hjc.todolist.util.PageQueryUtil;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface TaskMapper {
    int deleteByPrimaryKey(Long taskId);
    int deleteBatch(Long[] ids);

    int insert(Task task);

    int insertSelective(Task task);

    Task selectByTaskName(@Param("taskName") String taskName);

    int updateByPrimaryKeySelective(Task task);

    List<Task> findTaskList(PageQueryUtil pageUtil);
    Task findTaskById(Long taskId);

    int getTotalTasks(PageQueryUtil pageUtil);
}
